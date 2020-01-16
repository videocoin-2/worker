package graphite

import (
	"crypto/tls"
	"errors"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/videocoin/telegraf"
	tlsint "github.com/videocoin/telegraf/core/tls"
	"github.com/videocoin/telegraf/plugins/outputs"
	"github.com/videocoin/telegraf/plugins/serializers"
)

type Graphite struct {
	GraphiteTagSupport bool
	// URL is only for backwards compatibility
	Servers  []string
	Prefix   string
	Template string
	Timeout  int
	conns    []net.Conn
	tlsint.ClientConfig
}

func (g *Graphite) Connect() error {
	// Set default values
	if g.Timeout <= 0 {
		g.Timeout = 2
	}
	if len(g.Servers) == 0 {
		addr := os.Getenv("GRAPHITE_ADDR")
		if addr == "" {
			g.Servers = append(g.Servers, "localhost:2003")
		} else {
			g.Servers = append(g.Servers, addr)
		}
	}

	// Set tls config
	tlsConfig, err := g.ClientConfig.TLSConfig()
	if err != nil {
		return err
	}

	// Get Connections
	var conns []net.Conn
	for _, server := range g.Servers {
		// Dialer with timeout
		d := net.Dialer{Timeout: time.Duration(g.Timeout) * time.Second}

		// Get secure connection if tls config is set
		var conn net.Conn
		if tlsConfig != nil {
			conn, err = tls.DialWithDialer(&d, "tcp", server, tlsConfig)
		} else {
			conn, err = d.Dial("tcp", server)
		}

		if err == nil {
			conns = append(conns, conn)
		}
	}
	g.conns = conns
	return nil
}

func (g *Graphite) Close() error {
	// Closing all connections
	for _, conn := range g.conns {
		conn.Close()
	}
	return nil
}

// We need check eof as we can write to nothing without noticing anything is wrong
// the connection stays in a close_wait
// We can detect that by finding an eof
// if not for this, we can happily write and flush without getting errors (in Go) but getting RST tcp packets back (!)
// props to Tv via the authors of carbon-relay-ng` for this trick.
func checkEOF(conn net.Conn) {
	b := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	num, err := conn.Read(b)
	if err == io.EOF {
		log.Printf("E! Conn %s is closed. closing conn explicitly", conn)
		conn.Close()
		return
	}
	// just in case i misunderstand something or the remote behaves badly
	if num != 0 {
		log.Printf("I! conn %s .conn.Read data? did not expect that.  data: %s\n", conn, b[:num])
	}
	// Log non-timeout errors or close.
	if e, ok := err.(net.Error); !(ok && e.Timeout()) {
		log.Printf("E! conn %s checkEOF .conn.Read returned err != EOF, which is unexpected.  closing conn. error: %s\n", conn, err)
		conn.Close()
	}
}

// Choose a random server in the cluster to write to until a successful write
// occurs, logging each unsuccessful. If all servers fail, return error.
func (g *Graphite) Write(metrics []telegraf.Metric) error {
	// Prepare data
	var batch []byte
	s, err := serializers.NewGraphiteSerializer(g.Prefix, g.Template, g.GraphiteTagSupport)
	if err != nil {
		return err
	}

	for _, metric := range metrics {
		buf, err := s.Serialize(metric)
		if err != nil {
			log.Printf("E! Error serializing some metrics to graphite: %s", err.Error())
		}
		batch = append(batch, buf...)
	}

	err = g.send(batch)

	// try to reconnect and retry to send
	if err != nil {
		log.Println("E! Graphite: Reconnecting and retrying: ")
		g.Connect()
		err = g.send(batch)
	}

	return err
}

func (g *Graphite) send(batch []byte) error {
	// This will get set to nil if a successful write occurs
	err := errors.New("Could not write to any Graphite server in cluster\n")

	// Send data to a random server
	p := rand.Perm(len(g.conns))
	for _, n := range p {
		if g.Timeout > 0 {
			g.conns[n].SetWriteDeadline(time.Now().Add(time.Duration(g.Timeout) * time.Second))
		}
		checkEOF(g.conns[n])
		if _, e := g.conns[n].Write(batch); e != nil {
			// Error
			log.Println("E! Graphite Error: " + e.Error())
			// Close explicitly
			g.conns[n].Close()
			// Let's try the next one
		} else {
			// Success
			err = nil
			break
		}
	}

	return err
}

func init() {
	outputs.Add("graphite", func() telegraf.Output {
		return &Graphite{}
	})
}

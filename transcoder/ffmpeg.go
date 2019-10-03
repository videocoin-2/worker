package transcoder

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/armon/circbuf"
	v1 "github.com/videocoin/cloud-api/dispatcher/v1"
)

func (t *Transcoder) runFFmpeg(task *v1.Task, wg *sync.WaitGroup, errCh chan error) {
	stopCh := make(chan bool, 1)
	defer func() {
		close(errCh)
		wg.Done()
		stopCh <- true
		close(stopCh)
		t.logger.Debug("ffmpeg has been completed")
	}()

	t.logger.Debugf("starting ffmpeg")
	t.logger.Debugf("%s", task.Cmdline)

	stdoutPipe, err := t.cmd.StdoutPipe()
	if err != nil {
		fmtErr := fmt.Errorf("ffmpeg: %s", err)
		errCh <- fmtErr
		return
	}

	stderrPipe, err := t.cmd.StderrPipe()
	if err != nil {
		fmtErr := fmt.Errorf("ffmpeg: %s", err)
		errCh <- fmtErr
		return
	}

	err = t.cmd.Start()
	if err != nil {
		fmtErr := fmt.Errorf("ffmpeg: %s", err)
		errCh <- fmtErr
		return
	}

	stdouterr := bytes.NewBuffer(nil)
	ffmpegout, _ := circbuf.NewBuffer(1024 * 4)

	go io.Copy(stdouterr, stderrPipe)
	go io.Copy(stdouterr, stdoutPipe)

	go func(stopCh chan bool) {
		for {
			select {
			case <-stopCh:
				return
			default:
				buf := bytes.NewBuffer(nil)
				for {
					b, err := stdouterr.ReadByte()
					if err != nil {
						break
					}

					if b == '\x00' || b == 'x' {
						continue
					}

					if b == '\r' || b == '\n' {
						if ffmpegout != nil {
							ffmpegout.Write([]byte{'\n'})
						}

						line := buf.String()
						buf.Reset()

						if len(line) > 0 {
							t.logger.
								WithField("system", "ffmpeg").
								Debugf("ffmpeg: %s", line)
						}
					} else {
						if ffmpegout != nil {
							ffmpegout.Write([]byte{b})
						}
						buf.WriteByte(b)
					}
				}
			}
		}
	}(stopCh)

	t.logger.Info("waiting ffmpeg")

	err = t.cmd.Wait()
	if err != nil {
		fmtErr := err
		if ErrExitStatusInterrupt.Error() == err.Error() {
			t.logger.Warning("ffmpeg execution has been canceled")
		} else {
			fmtErr = fmt.Errorf("ffmpeg: %s", err)
		}

		errCh <- fmtErr
		return
	}
}

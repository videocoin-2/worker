package transcode

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"log"

	bc "github.com/VideoCoin/common/bcops"
	pb "github.com/VideoCoin/common/proto"
	"github.com/VideoCoin/common/streamManager"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// New initialize and return a new Service object
func New() (*Service, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := LoadConfig()

	// Generate unique connection name
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}

	managerConn, err := grpc.Dial(cfg.ManagerRPCADDR, opts...)
	if err != nil {
		panic(err)
	}

	manager := pb.NewManagerServiceClient(managerConn)
	fmt.Printf("%+v", manager)
	status, err := manager.Health(context.Background(), &empty.Empty{})
	if status.GetStatus() != "healthy" || err != nil {
		return nil, fmt.Errorf("failed to get healthy status: %v", err)
	}

	ctx := context.Background()

	client, err := ethclient.Dial(cfg.BlockchainURL)
	if err != nil {
		panic(err)
	}

	managerAddress := common.HexToAddress(cfg.SMCA)

	sm, err := streamManager.NewManager(managerAddress, client)
	if err != nil {
		panic(err)
	}

	key, err := bc.LoadBcPrivKeys(cfg.KeyFile, cfg.Password)
	if err != nil {
		panic(err)
	}

	bcAuth, err := bc.GetBCAuth(client, key)
	if err != nil {
		panic(err)
	}

	return &Service{
		sm:       sm,
		bcAuth:   bcAuth,
		bcClient: client,
		cfg:      cfg,
		manager:  manager,
		ctx:      ctx,
		log:      logrus.WithField("name", "transcoder"),
	}, nil

}

func (s *Service) reportStatus(userID string, streamID int64, status string) error {
	ctx := context.Background()
	_, err := s.manager.UpdateStreamStatus(ctx, &pb.UpdateStreamStatusRequest{
		StreamId: streamID,
		Status:   status,
	})

	if err != nil {
		return err
	}

	return nil
}

// Start creates new service and blocks until stop signal
func Start() error {
	s, err := New()
	if err != nil {
		panic(err)
	}

	task, err := s.manager.GetJob(s.ctx, &pb.GetJobRequest{})
	if err != nil {
		return err
	}

	task.Status = pb.WorkOrderStatusTranscoding.String()

	_, err = s.manager.UpdateStreamStatus(s.ctx, &pb.UpdateStreamStatusRequest{
		StreamId: task.StreamId,
		Status:   task.Status,
	})

	if err != nil {
		s.log.Errorf("failed to update stream status: %s", err.Error())
	}

	return s.handleTranscodeTask(task)

}

func (s *Service) handleTranscodeTask(workOrder *pb.WorkOrder) error {

	s.log.Infof("starting transcode task: %d using input: %s with stream_id: %d", workOrder.Id, workOrder.InputUrl, workOrder.StreamId)

	dir := path.Join(s.cfg.OutputDir, fmt.Sprintf("%d", workOrder.StreamId))
	m3u8 := path.Join(dir, "index.m3u8")

	for _, b := range bitrates {
		fullDir := fmt.Sprintf("%s/%d", dir, b)
		err := prepareDir(fullDir)
		if err != nil {
			s.log.Errorf("failed to prepare directory [ %s ]: %s", fullDir, err.Error())
		}

		s.log.Infof("monitoring chunks in %s", fullDir)

		go s.monitorChunks(fullDir, workOrder)
		go s.SyncDir(workOrder, fullDir, b)
	}

	if err := s.generatePlaylist(workOrder.StreamId, m3u8); err != nil {
		s.log.Fatalf("failed to generate playlist: %s", err.Error())
	}

	args := buildCmd(workOrder.InputUrl, dir)

	s.transcode(args, workOrder.InputUrl)

	return nil
}

func (s *Service) monitorChunks(dir string, task *pb.WorkOrder) {
	for {
		time.Sleep(2 * time.Second)
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Printf("failed to read dir: %s", err.Error())
		}

		if len(files) < 2 {
			continue
		}

		break
	}

	task.Status = pb.WorkOrderStatusReady.String()
	update := &pb.UpdateStreamStatusRequest{
		StreamId: task.StreamId,
		Status:   task.Status,
	}
	_, err := s.manager.UpdateStreamStatus(s.ctx, update)

	if err != nil {
		s.log.Errorf("failed to update stream status: %s", err.Error())
	}
	fmt.Println(task.Status, task.StreamId)

}

func (s *Service) transcode(args []string, streamurl string) {
	waitForStreamReady(streamurl)
	log.Println("starting transcode")
	out, err := exec.Command("ffmpeg", args...).CombinedOutput()
	if err != nil {
		s.log.Fatalf("failed to transcode: err : %s output: %s", err.Error(), string(out))
	}

	s.log.Info("transcode complete")
}

func (s *Service) generatePlaylist(streamID int64, filename string) error {
	m3u8 := []byte(fmt.Sprintf(`#EXTM3U
#EXT-X-VERSION:6
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1048576,RESOLUTION=640x360,CODECS="avc1.42e00a,mp4a.40.2"
%d/index.m3u8
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=3145728,RESOLUTION=842x480,CODECS="avc1.42e00a,mp4a.40.2"
%d/index.m3u8
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=5242880,RESOLUTION=1280x720,CODECS="avc1.42e00a,mp4a.40.2"
%d/index.m3u8
`, bitrates[0], bitrates[1], bitrates[2]))

	err := ioutil.WriteFile(filename, m3u8, 0755)
	if err != nil {
		s.log.Errorf("failed to write file: %s", err.Error())
		return err
	}

	reader := bytes.NewReader(m3u8)

	err = s.Upload(fmt.Sprintf("%d/%s", streamID, "index.m3u8"), reader)
	if err != nil {
		s.log.Errorf("failed to upload: %s", err.Error())
		return err
	}

	return nil
}

func waitForStreamReady(streamurl string) {
	maxretry := 10
	for i := 0; i < maxretry; i++ {
		resp, _ := http.Head(streamurl)
		if resp.StatusCode == 200 {
			return
		}
		log.Printf("waiting for stream %s to become ready...", streamurl)
		time.Sleep(30 * time.Second)
	}
}

func prepareDir(dir string) error {
	return os.MkdirAll(dir, 0777)
}

func buildCmd(inputURL string, dir string) []string {
	cmd := []string{"-re", "-i", inputURL}

	for _, b := range bitrates {
		args := fmt.Sprintf("-hls_allow_cache 0 -hls_flags append_list -f ssegment -b:v %d -strict -2 -c:v h264 -profile:v main -segment_list_flags live -segment_time 10 -segment_format mpegts -an -segment_list %s/%d/index.m3u8 %s/%d/%%d.ts", b, dir, b, dir, b)
		cmd = append(cmd, strings.Split(args, " ")...)

	}

	return cmd

}

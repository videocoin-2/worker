package service

import (
	"github.com/google/uuid"
	dispatcherv1 "github.com/videocoin/cloud-api/dispatcher/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/transcode/telegraf"
	"github.com/videocoin/transcode/transcoder"
	"google.golang.org/grpc"
)

type Service struct {
	cfg        *Config
	dispatcher dispatcherv1.DispatcherServiceClient
	transcoder *transcoder.Transcoder

	MachineID string
}

func NewService(cfg *Config) (*Service, error) {
	dlogger := cfg.Logger.WithField("system", "dispatchercli")
	dGrpcDialOpts := grpcutil.ClientDialOptsWithRetry(dlogger)
	dispatcherConn, err := grpc.Dial(cfg.DispatcherRPCAddr, dGrpcDialOpts...)
	if err != nil {
		return nil, err
	}
	dispatcher := dispatcherv1.NewDispatcherServiceClient(dispatcherConn)

	trans, err := transcoder.NewTranscoder(cfg.Logger.WithField("system", "transcoder"))
	if err != nil {
		return nil, err
	}

	machineID := uuid.New()

	svc := &Service{
		cfg:        cfg,
		dispatcher: dispatcher,
		transcoder: trans,
		MachineID:  machineID.String(),
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.transcoder.Start()
	go telegraf.Run(
		s.cfg.Logger.WithField("system", "telegraf"),
		s.MachineID)

	return nil
}

func (s *Service) Stop() error {
	s.transcoder.Stop()
	return nil
}

// func init() {
// 	logrus.SetFormatter(&logrus.JSONFormatter{})
// 	mRand.Seed(time.Now().Unix())
// }

// // New initialize and return a new Service object
// func newService() (*Service, error) {
// 	log := logrus.WithField("service", "transcode")
// 	cfg := LoadConfig()

// 	level, _ := logrus.ParseLevel(cfg.LogLevel)
// 	logrus.SetLevel(level)
// 	// Generate unique connection name
// 	b := make([]byte, 8)
// 	rand.Read(b)

// 	opts := []grpc.DialOption{grpc.WithInsecure()}

// 	managerConn, err := grpc.Dial(cfg.ManagerRPCADDR, opts...)
// 	if err != nil {
// 		log.Fatalf("failed to dial manager: %s", err.Error())
// 	}

// 	verifierConn, err := grpc.Dial(cfg.VerifierRPCADDR, opts...)
// 	if err != nil {
// 		log.Fatalf("failed to dial verifier: %s", err.Error())
// 	}

// 	manager := manager_v1.NewManagerServiceClient(managerConn)
// 	status, err := manager.Health(context.Background(), &types.Empty{})
// 	if status.GetStatus() != "healthy" || err != nil {
// 		return nil, fmt.Errorf("failed to get healthy status from manager")
// 	}

// 	v := verifier_v1.NewVerifierServiceClient(verifierConn)
// 	status, err = v.Health(context.Background(), &types.Empty{})
// 	if status.GetStatus() != "healthy" || err != nil {
// 		return nil, fmt.Errorf("failed to get healthy status from verifier")
// 	}

// 	ctx := context.Background()

// 	return &Service{
// 		cfg:      cfg,
// 		manager:  manager,
// 		verifier: v,
// 		ctx:      ctx,
// 		log:      log,
// 	}, nil

// }

// // Start creates new service and blocks until stop signal
// func Start() error {
// 	s, err := newService()
// 	if err != nil {
// 		return err
// 	}

// 	uid, err := uuid4.New()
// 	if err != nil {
// 		return err
// 	}

// 	s.register(uid)

// 	s.log.Infof("transcoder_id: %s", uid)

// 	go s.handleExit()

// 	s.pollForWork()

// 	return nil
// }

// func (s *Service) register(uid string) {

// 	var (
// 		cores    int32
// 		mhz      float64
// 		memtotal uint64
// 	)

// 	info, err := cpu.Info()
// 	if err != nil {
// 		s.log.Errorf("failed to get cpu info: %s", err.Error())
// 	} else {
// 		cores = info[0].Cores
// 		mhz = info[0].Mhz
// 	}

// 	memInfo, err := mem.VirtualMemory()
// 	if err != nil {
// 		s.log.Errorf("failed to get memory info: %s", err.Error())
// 	} else {
// 		memtotal = memInfo.Total
// 	}

// 	_, err = s.manager.RegisterTranscoder(context.Background(), &transcoder_v1.Transcoder{
// 		Id:          uid,
// 		CpuCores:    cores,
// 		CpuMhz:      mhz,
// 		TotalMemory: memtotal,
// 		Status:      transcoder_v1.TranscoderStatusAvailable,
// 	})

// 	if err != nil {
// 		s.log.Errorf("failed to register transcoder: %s", err.Error())
// 	}
// }

// var gracefulStop = make(chan os.Signal)

// func (s *Service) pollForWork() {
// 	ticker := time.NewTicker(5 * time.Second)
// 	for range ticker.C {
// 		log.Print("polling for work...")

// 		assignment, err := s.manager.GetWork(context.Background(), &types.Empty{})
// 		if err != nil {
// 			continue
// 		}

// 		s.eth = NewEth(s.cfg)
// 		s.eth.connect(cfg.BlockchainURL, cfg.SMCA, assignment.Job.StreamAddress)

// 		log.Printf("work found: id=%s", assignment.Job.Id)
// 		log.Printf("using key: %s", s.eth.kv.Key)
// 		log.Printf("value: %s", string(s.eth.kv.Value))

// 		if s.handleTranscode(assignment) != nil {
// 			os.Exit(0)
// 		}
// 	}
// }

// func (s *Service) handleExit() {
// 	signal.Notify(gracefulStop, syscall.SIGTERM)
// 	signal.Notify(gracefulStop, syscall.SIGINT)
// 	signal.Notify(gracefulStop, syscall.Signal(0))

// 	sig := <-gracefulStop

// 	s.log.Infof("exiting: %s", sig.String())

// }

// func (s *Service) handleTranscode(a *transcoder_v1.Assignment) error {
// 	dir := path.Join(s.cfg.OutputDir, a.Job.Id)
// 	m3u8 := path.Join(dir, "index.m3u8")

// 	cmd := buildCmd(a.Job.TranscodeInputUrl, dir, a.Profile)
// 	var stopChan = make(chan bool)

// 	fullDir := fmt.Sprintf("%s/%d", dir, a.Profile.Bitrate)
// 	err := prepareDir(fullDir)

// 	if err != nil {
// 		return err
// 	}

// 	go s.syncDir(stopChan, cmd, a.Job, fullDir, a.Profile.Bitrate)

// 	if err := s.generatePlaylist(a.Job.Id, m3u8, a.Profile.Bitrate); err != nil {
// 		return err
// 	}

// 	s.transcode(cmd,
// 		stopChan,
// 		a.Job.TranscodeInputUrl,
// 	)

// 	return nil
// }

// func (s *Service) transcode(
// 	cmd *exec.Cmd,
// 	stop chan bool,
// 	streamurl string,
// ) {
// 	s.waitForStreamReady(streamurl)

// 	out, err := cmd.CombinedOutput()

// 	// Allow any chunks being uploaded to finish
// 	time.Sleep(20 * time.Second)

// 	if err != nil {
// 		s.log.Errorf("failed to transcode: err : %s output: %s",
// 			err.Error(), string(out),
// 		)
// 		cmd.Process.Kill()
// 		os.Exit(0)
// 	}

// 	stop <- true

// 	s.log.Info("transcode complete")

// }

// func (s *Service) waitForStreamReady(streamurl string) {
// 	maxretry := 15
// 	for i := 0; i < maxretry; i++ {
// 		resp, err := http.Head(streamurl)
// 		if err != nil {
// 			s.log.Errorf("failed to request stream: %s", err.Error())
// 		} else if resp.StatusCode == 200 {
// 			return
// 		}
// 		s.log.Infof("waiting for stream %s to become ready...", streamurl)
// 		time.Sleep(2 * time.Second)
// 	}

// }

// func prepareDir(dir string) error {
// 	return os.MkdirAll(dir, 0777)
// }

// func buildCmd(inputURL string, dir string, profile *profiles_v1.Profile) *exec.Cmd {
// 	process := []string{"-re", "-i", inputURL}

// 	args := fmt.Sprintf("-max_muxing_queue_size 9999 -live_start_index -1 -b:v %d -vf scale=%d:-2 -strict -2 -c:v libx264 -c:a aac -bsf:v h264_mp4toannexb -map 0 -f segment -segment_time 2 -segment_format mpegts -segment_list %s/%d/index.m3u8 -segment_list_type m3u8 %s/%d/%%d.ts", profile.Bitrate, profile.Width, dir, profile.Bitrate, dir, profile.Bitrate)
// 	process = append(process, strings.Split(args, " ")...)

// 	cmd := exec.Command("ffmpeg", process...)

// 	return cmd

// }

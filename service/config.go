package service

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// Config default config for transcoder
type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	DispatcherRPCAddr  string `required:"true" envconfig:"DISPATCHER_ADDR" default:"d.dev.videocoin.network:5008"`
	RPCNodeURL         string `required:"true" envconfig:"RPC_NODE_URL" default:"https://dev1:D6msEL93LJT5RaPk@rpc.dev.kili.videocoin.network"`
	SyncerURL          string `required:"true" envconfig:"SYNCER_URL" default:"https://dev.videocoin.network/api/v1/sync"`
	OutputDir          string `required:"true" envconfig:"OUTPUT_DIR" default:"/tmp"`
	StakingManagerAddr string `required:"true" envconfig:"STAKING_MANAGER_ADDR" default:"0x817ec8E65252E80dB27eFbBceE940AD917AC78FF"`

	ClientID string `envconfig:"CLIENT_ID"`
	Key      string `envconfig:"KEY"`
	Secret   string `envconfig:"SECRET"`

	Internal bool `envconfig:"INTERNAL"`
}

var cfg Config
var once sync.Once

func LoadConfig() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatal(err.Error())
		}

	})
	return &cfg
}

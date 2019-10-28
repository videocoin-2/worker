package service

import (
	"github.com/sirupsen/logrus"
)

// Config default config for transcoder
type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	DispatcherRPCAddr string `required:"true" envconfig:"DISPATCHER_ADDR" default:"d.snb.videocoin.network:5008"`
	OutputDir         string `required:"true" envconfig:"OUTPUT_DIR" default:"/tmp" description:"local folder for ts chunks"`
	BlockchainURL     string `required:"true" envconfig:"BLOCKCHAIN_URL" default:"http://admin:VideoCoinS3cr3t@rpc.dev.videocoin.network"`
	SMCA              string `required:"true" envconfig:"SMCA" default:"0xEa91ac0B88F84e91e79Caa871d2EB04eF5133721" description:"stream manager contract address"`
	Key               string `required:"true" envconfig:"KEY"`
	Secret            string `required:"true" envconfig:"SECRET" default:"transcoder"`
	ClientID          string `required:"true" envconfig:"CLIENT_ID"`
}

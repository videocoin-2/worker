module github.com/videocoin/worker

go 1.14

require (
	cloud.google.com/go v0.51.0 // indirect
	github.com/AlekSi/pointer v1.1.0
	github.com/armon/circbuf v0.0.0-20150827004946-bbbad097214e // indirect
	github.com/ethereum/go-ethereum v1.9.12
	github.com/evalphobia/logrus_sentry v0.8.2 // indirect
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/grafov/m3u8 v0.11.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.1.10
	github.com/shirou/gopsutil v2.20.2+incompatible
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v0.0.7
	github.com/videocoin/cloud-api v0.3.0
	github.com/videocoin/cloud-pkg v0.0.7
	github.com/videocoin/go-bridge v0.0.2 // indirect
	github.com/videocoin/go-protocol v0.0.6
	github.com/videocoin/go-staking v0.0.0-20200410180201-6944f4d9a28b
	github.com/videocoin/oauth2 v0.0.0-20200430234055-8f1bc1d0e599
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/api v0.20.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/grpc v1.27.1
)

replace github.com/videocoin/cloud-pkg => ../cloud-pkg

replace github.com/videocoin/cloud-api => ../cloud-api

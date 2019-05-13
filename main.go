package main

import (
	"github.com/rebornwwp/easyproxy/config"
	gw "github.com/rebornwwp/easyproxy/gateway"
	"github.com/rebornwwp/easyproxy/log"
	"github.com/rebornwwp/easyproxy/util"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
)

const (
	DefaultConfigFile = "config.json"
	DefaultLogFile    = "proxy.log"
)

type EasyServer struct {
	proxyServer *gw.ProxyServer
}

func CreateEasyServer() *EasyServer {
	return &EasyServer{proxyServer: new(gw.ProxyServer)}
}

func (server *EasyServer) Init(config *config.Config) {
	server.proxyServer.Init(config)
}

func (server *EasyServer) Start() {
	server.proxyServer.Start()
}

func (server *EasyServer) CatchStopSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-sig
		server.Stop()
	}()
}

func (server *EasyServer) Stop() {
	server.proxyServer.Stop()
}
func main() {
	log.Init(DefaultLogFile)

	homePath := util.HomePath()
	conf, err := config.NewConfig(path.Join(homePath, DefaultConfigFile))
	if err == nil {
		runtime.GOMAXPROCS(conf.MaxProcessor)
		server := CreateEasyServer()
		server.Init(conf)
		server.CatchStopSignal()
		server.Start()
	}
}

package gateway

import (
	"log"
	"net"
	"time"

	"github.com/rebornwwp/easyproxy/config"
	"github.com/rebornwwp/easyproxy/proxy"
	"github.com/rebornwwp/easyproxy/util"
)

type ProxyServer struct {
	Host     string
	Port     uint16
	BeatTime int
	Listener net.Listener
	Proxy    proxy.Proxy
	on       bool
}

func (proxyServer *ProxyServer) Init(config *config.Config) {
	proxyServer.on = false
	proxyServer.Host = config.Host
	proxyServer.Port = config.Port
	proxyServer.BeatTime = config.HeartBeat
	proxyServer.setProxy(config)
}

func (proxyServer *ProxyServer) setProxy(config *config.Config) {
	proxyServer.Proxy = new(proxy.EasyProxy)
	proxyServer.Proxy.Init(config)
}

func (proxyServer *ProxyServer) Address() string {
	return util.HostPortToAddress(proxyServer.Host, proxyServer.Port)
}

func (proxyServer *ProxyServer) Start() {
	local, err := net.Listen("tcp", proxyServer.Address())
	if err != nil {
		log.Panic("proxy server start err", err)
	}
	log.Println("proxy server start ok")
	proxyServer.on = true
	proxyServer.Listener = local
	proxyServer.heartBeat()
	for proxyServer.on {
		conn, err := proxyServer.Listener.Accept()
		if err == nil {
			go proxyServer.Proxy.Dispatch(conn)
		} else {
			log.Println("client connect to server error:", err)
		}
	}
	proxyServer.Proxy.Close()
}

func (proxyServer *ProxyServer) heartBeat() {
	ticker := time.NewTicker(time.Second * time.Duration(proxyServer.BeatTime))
	go func() {
		for {
			select {
			case <-ticker.C:
				proxyServer.Proxy.Check()
			}
		}
	}()
}

func (proxyServer *ProxyServer) Stop() {
	proxyServer.Listener.Close()
	proxyServer.Proxy.Close()
	proxyServer.on = false
	log.Println("easy proxy server stop ok")
}

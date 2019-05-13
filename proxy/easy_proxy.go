package proxy

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/rebornwwp/easyproxy/config"
	"github.com/rebornwwp/easyproxy/proxy/schedule"
	"github.com/rebornwwp/easyproxy/structure"
)

const (
	DefaultTimeOut = 3
)

type EasyProxy struct {
	data     *ProxyData
	strategy schedule.Strategy
}

func (proxy *EasyProxy) Init(config *config.Config) {
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
	proxy.setStrategy(config.Strategy)
	InitStatistic(proxy.data)
}

func (proxy *EasyProxy) setStrategy(name string) {
	proxy.strategy = schedule.GetStrategy(name)
	proxy.strategy.Init()
}

func (proxy *EasyProxy) Check() {
	for _, backend := range proxy.data.Backends {
		_, err := net.Dial("tcp", backend.Url())
		if err != nil {
			proxy.Clean(backend.Url())
		}
	}
	for _, dead := range proxy.data.Deads {
		_, err := net.Dial("tcp", dead.Url())
		if err == nil {
			proxy.Recover(dead.Url())
		}
	}
}

func (proxy *EasyProxy) Dispatch(conn net.Conn) {
	if proxy.isBackendAvailable() {
		servers := proxy.data.BackendUrls()
		url := proxy.strategy.Choose(conn.RemoteAddr().String(), servers)
		proxy.transfer(conn, url)
	} else {
		conn.Close()
		log.Println("no backend available now,please check your server!")
	}
}

func (proxy *EasyProxy) transfer(local net.Conn, remote string) {
	remoteConn, err := net.DialTimeout("tcp", remote, DefaultTimeOut*time.Second)
	if err != nil {
		local.Close()
		proxy.Clean(remote)
		log.Println("connect backend err", err)
		return
	}
	sync := make(chan int)
	channel := &structure.Channel{Src: local, Dst: remoteConn}
	go proxy.putChannel(channel)
	go proxy.safeCopy(local, remoteConn, sync)
	go proxy.safeCopy(remoteConn, local, sync)
	go proxy.closeChannel(channel, sync)
}

func (proxy *EasyProxy) putChannel(channel *structure.Channel) {
	proxy.data.ChannelManager.PutChannel(channel)
}

func (proxy *EasyProxy) safeCopy(from net.Conn, to net.Conn, sync chan<- int) {
	io.Copy(from, to)
	defer from.Close()
	sync <- 1
}

func (proxy *EasyProxy) closeChannel(chanel *structure.Channel, sync <-chan int) {
	for i := 0; i < structure.ChannelPairNum; i++ {
		<-sync
	}
	proxy.data.ChannelManager.DeleteChannel(chanel)
}

func (proxy *EasyProxy) isBackendAvailable() bool {
	return len(proxy.data.Backends) > 0
}

func (proxy *EasyProxy) Clean(url string) {
	proxy.data.deleteBackend(url)
}

func (proxy *EasyProxy) Recover(url string) {
	proxy.data.deleteDeads(url)
}

func (proxy *EasyProxy) Close() {
	proxy.data.Clean()
}

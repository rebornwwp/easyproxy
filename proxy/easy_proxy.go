package proxy

import (
	"github.com/rebornwwp/easyproxy/config"
	"github.com/rebornwwp/easyproxy/proxy/schedule"
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

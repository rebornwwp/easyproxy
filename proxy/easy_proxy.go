package proxy

import "github.com/rebornwwp/easyproxy/proxy/schedule"

const (
	DefaultTimeOut = 3
)

type EasyProxy struct {
	proxyData *ProxyData
	strategy  schedule.Strategy
}

package proxy

import (
	"sync"

	"github.com/rebornwwp/easyproxy/config"
	"github.com/rebornwwp/easyproxy/structure"
)

type ProxyData struct {
	Service        string
	Host           string
	Port           uint16
	Backends       map[string]structure.Backend
	Deads          map[string]structure.Backend
	ChannelManager *structure.ChannelManager
	mutex          *sync.RWMutex
}

func (proxyData *ProxyData) Init(config *config.Config) {
	proxyData.Service = config.Service
	proxyData.Host = config.Host
	proxyData.Port = config.Port
	proxyData.ChannelManager = new(structure.ChannelManager)
	proxyData.ChannelManager.Init()
	proxyData.InitBackends(config.Backends)
	proxyData.InitDeads()
	proxyData.mutex = new(sync.RWMutex)
}

func (proxyData *ProxyData) InitBackends(backends []structure.Backend) {
	proxyData.Backends = make(map[string]structure.Backend)
	for _, backend := range backends {
		proxyData.Backends[backend.Url()] = backend
	}
}

func (proxyData *ProxyData) InitDeads() {
	proxyData.Deads = make(map[string]structure.Backend)
}

func (proxyData *ProxyData) BackendUrls() []string {
	proxyData.mutex.RLock()
	_map := proxyData.Backends
	proxyData.mutex.RUnlock()
	urls := make([]string, 0, len(_map))
	for url, _ := range _map {
		urls = append(urls, url)
	}
	return urls
}

func (proxyData *ProxyData) deleteBackend(url string) {
	proxyData.mutex.Lock()
	defer proxyData.mutex.Unlock()
	if backend, ok := proxyData.Backends[url]; ok {
		proxyData.Deads[url] = backend
		delete(proxyData.Backends, url)
	}
}

func (proxyData *ProxyData) deleteDeads(url string) {
	proxyData.mutex.Lock()
	defer proxyData.mutex.Unlock()
	if backend, ok := proxyData.Deads[url]; ok {
		proxyData.Backends[url] = backend
		delete(proxyData.Deads, url)
	}
}

func clean(_map map[string]structure.Backend) {
	for url := range _map {
		delete(_map, url)
	}
}

func (proxyData *ProxyData) Clean() {
	clean(proxyData.Backends)
	clean(proxyData.Deads)
	proxyData.ChannelManager.Clean()
}

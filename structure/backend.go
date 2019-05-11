package structure

import "github.com/rebornwwp/easyproxy/util"

// Backend 代表一个后台的基本信息
type Backend struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

// Url 获取backend的url
func (backend Backend) Url() string {
	return util.HostPortToAddress(backend.Host, backend.Port)
}

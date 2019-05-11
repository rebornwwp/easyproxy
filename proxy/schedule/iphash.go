package schedule

import "github.com/rebornwwp/easyproxy/util"

type IpHash struct{}

func (strategy *IpHash) Init() {}

func (strategy *IpHash) Choose(client string, servers []string) string {
	ip := util.UrlToHost(client)
	ipInt := util.IpToInt(ip)
	length := len(servers)
	return servers[ipInt%length]
}

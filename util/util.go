package util

import "strconv"

// HostPortToAddress form host port to "host:port"
func HostPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}

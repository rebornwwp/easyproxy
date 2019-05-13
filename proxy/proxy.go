package proxy

import (
	"github.com/rebornwwp/easyproxy/config"
	"net"
)

type Proxy interface {
	Init(*config.Config)
	Check()
	Clean(string)
	Recover(string)
	Dispatch(conn net.Conn)
	Close()
}

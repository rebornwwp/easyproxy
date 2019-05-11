package structure

import "net"

const ChannelPairNum = 2

type Channel struct {
	Src net.Conn
	Dst net.Conn
}

func (channel *Channel) SrcUrl() string {
	return channel.Src.RemoteAddr().String()
}

func (channel *Channel) DstUrl() string {
	return channel.Dst.RemoteAddr().String()
}

func (channel *Channel) Close() {
	_ := channel.Dst.Close()
	_ := channel.Src.Close()
}

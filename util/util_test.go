package util

import "testing"

func TestHostPortToAddress(t *testing.T) {
	host := "123.10.123.24"
	var port uint16 = 1234
	t.Log(HostPortToAddress(host, port))
}

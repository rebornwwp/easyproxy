package structure

import "testing"

func TestUrl(t *testing.T) {
	backend := Backend{
		Host: "123.123.123.1",
		Port: uint16(1234),
	}
	t.Log(backend.Url())
}

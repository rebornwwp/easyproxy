package util

import (
	"reflect"
	"strconv"
	"strings"
)

// HostPortToAddress form host port to "host:port"
func HostPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}

func UrlToHost(address string) string {
	return strings.Split(address, ":")[0]
}

func IpToInt(ip string) int {
	nums := strings.Split(ip, ".")
	ans := 0
	for _, num := range nums {
		ans = ans << 8
		tmp, _ := strconv.Atoi(num)
		ans += tmp
	}
	return ans
}

func SliceIndex(slice interface{}, element interface{}) int {
	index := -1
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		return index
	}
	ev := reflect.ValueOf(element).Interface()
	length := sv.Len()
	for i := 0; i < length; i++ {
		iv := sv.Index(i).Interface()
		if reflect.DeepEqual(iv, ev) {
			index = i
			break
		}
	}
	return index
}

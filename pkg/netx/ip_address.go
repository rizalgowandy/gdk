package netx

import (
	"net"
)

func GetIPv4() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return err.Error()
	}

	for _, address := range addresses {
		current, ok := address.(*net.IPNet)
		if !ok {
			continue
		}

		if current.IP.IsLoopback() {
			continue
		}

		if current.IP.To4() != nil {
			return current.IP.To4().String()
		}
	}

	return "ip v4: unavailable"
}

func GetIPv16() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return err.Error()
	}

	for _, address := range addresses {
		current, ok := address.(*net.IPNet)
		if !ok {
			continue
		}

		if current.IP.IsLoopback() {
			continue
		}

		if current.IP.To16() != nil {
			return current.IP.To16().String()
		}
	}

	return "ip v16: unavailable"
}

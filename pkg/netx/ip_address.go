package netx

import (
	"net"

	"github.com/rizalgowandy/gdk/pkg/syncx"
)

var (
	onceGetIPv4    syncx.Once
	onceGEtIPv4Res string

	onceGetIPv16    syncx.Once
	onceGEtIPv16Res string
)

func GetIPv4() string {
	onceGetIPv4.Do(func() {
		addresses, err := net.InterfaceAddrs()
		if err != nil {
			onceGEtIPv4Res = err.Error()
			return
		}

		for _, address := range addresses {
			current, ok := address.(*net.IPNet)
			if !ok {
				continue
			}

			if current.IP.IsLoopback() || current.IP.To4() == nil {
				continue
			}

			onceGEtIPv4Res = current.IP.To4().String()
			return
		}

		onceGEtIPv4Res = "ip v4: unavailable"
	})
	return onceGEtIPv4Res
}

func GetIPv16() string {
	onceGetIPv16.Do(func() {
		addresses, err := net.InterfaceAddrs()
		if err != nil {
			onceGEtIPv16Res = err.Error()
			return
		}

		for _, address := range addresses {
			current, ok := address.(*net.IPNet)
			if !ok {
				continue
			}

			if current.IP.IsLoopback() || current.IP.To16() == nil {
				continue
			}

			onceGEtIPv16Res = current.IP.To16().String()
			return
		}

		onceGEtIPv16Res = "ip v16: unavailable"
	})
	return onceGEtIPv16Res
}

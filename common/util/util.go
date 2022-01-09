package util

import (
	"github.com/google/uuid"
	"net"
	"strings"
)

// UUID v4
func GetUUID() string {
	uuidv4 := uuid.New().String()
	newuuid := strings.Replace(uuidv4, "-", "", -1)
	return newuuid
}

// 获取本机网卡IP
func GetLocalIP() (ipv4 string) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet
		isIpNet bool
		err     error
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}

	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址(有可能是Unix,socket地址)
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}

func GetAppFullName(name string) string {
	return strings.ToUpper(name) + "-" + GetUUID()
}

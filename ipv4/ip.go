/*
 * @Author: your name
 * @Date: 2021-06-07 21:39:45
 * @LastEditTime: 2021-06-07 22:00:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \utility\ipv4\ip.go
 */
package ipv4

import (
	"errors"
	"net"
	"strings"
)

/**
 * @description: 获取本地IPV4地址
 * @param {*}
 * @return {*}
 */
func LocalIPV4() (ip net.IP, err error) {
	ips, err := GetIPs()
	if err != nil {
		return nil, err
	}
	for _, ip := range ips {
		v4 := ip.To4()
		if v4 == nil || v4[0] == 127 { // loopback address
			continue
		}
		return v4, nil
	}
	return nil, errors.New("无法获取到IP地址")
}

/**
 * @description:获取所有IP地址
 * @param {*}
 * @return {*}
 */
func GetIPs() ([]net.IP, error) {
	var ips []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalUnicast() {
			ips = append(ips, ipnet.IP)
		}
	}
	return ips, nil
}

/**
 * @description: 获取出口本地地址
 * @param {*}
 * @return {*}
 */
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

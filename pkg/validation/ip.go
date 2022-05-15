package validation

import (
	"net"
	"strings"
)

// IsIPAddr return true if string ip contains a valid representation of an IPv4 or IPv6 address
func IsIPAddr(ip string) bool {
	ipaddr := net.ParseIP(NormaliseIPAddr(ip))
	return ipaddr != nil
}

// NormaliseIPAddr return ip adresse without /32 (IPv4 or /128 (IPv6)
func NormaliseIPAddr(ip string) string {
	if strings.HasSuffix(ip, "/32") && strings.Contains(ip, ".") { // single host (IPv4)
		ip = strings.TrimSuffix(ip, "/32")
	} else {
		if strings.HasSuffix(ip, "/128") { // single host (IPv6)
			ip = strings.TrimSuffix(ip, "/128")
		}
	}

	return ip
}

package netip

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func maskRange(ip net.IP, mask net.IPMask) (btm, top net.IP) {
	btm = ip.Mask(mask)
	top = make(net.IP, len(ip))
	copy(top, ip)
	for i, b := range mask {
		top[i] |= ^b
	}
	return
}

func ParseMask(address, netmask string) (string, error) {
	ip := net.ParseIP(address)
	nmip := net.ParseIP(netmask)
	if ip == nil || nmip == nil {
		return "", errors.New(fmt.Sprintf("either %s or %s couldn't be parsed as an IP",
			address, netmask))
	}
	// this is a bit of a hack, because using ParseIP to parse
	// something that's actually a v4 netmask doesn't quite work
	nm := net.IPMask(nmip.To4())
	cidr, bits := nm.Size()
	if ip.To4() != nil && nm != nil {
		if bits != 32 {
			return "", errors.New(fmt.Sprintf("%s doesn't look like a valid IPv4 netmask", netmask))
		}
	} else {
		// IPv6, hopefully
		nm = net.IPMask(nmip)
		cidr, bits = nm.Size()
		if bits != 128 {
			return "", errors.New(fmt.Sprintf("%s doesn't look like a valid IPv6 netmask", netmask))
		}
	}
	btm, top := maskRange(ip, nm)
	return fmt.Sprintf("%s/%d is in the range %s-%s and has the netmask %s",
		ip, cidr, btm, top, nmip), nil
}

// GetIPSubnet GetIPSubnet("192.168.15.1", "255.255.255.0")  =>  192.168.15.0/24 <nil>
func GetIPSubnet(ip, netmask string) (ipPrefix, prefix string, err error) {
	// Check ip
	if net.ParseIP(ip) == nil {
		return "", "", fmt.Errorf("invalid IP address %s", ip)
	}
	// Check netmask
	maskIP := net.ParseIP(netmask).To4()
	if maskIP == nil {
		return "", "", fmt.Errorf("invalid Netmask %s", netmask)
	}
	// Get prefix
	mask := net.IPMask(maskIP)
	p, _ := mask.Size()

	// Get network
	sPrefix := strconv.Itoa(p)
	_, network, err := net.ParseCIDR(ip + "/" + sPrefix)
	if err != nil {
		return "", "", err
	}
	return network.IP.String() + "/" + sPrefix, sPrefix, nil
}

// SubnetString SubnetString("24") => "255.255.255.0"
func SubnetString(cidr string) (string, error) {
	var maskList []string
	var netMask string
	cidrInt, err := strconv.ParseUint(cidr, 10, 8)
	if err != nil {
		return "", err
	}
	for i := 0; i < 4; i++ {
		tmp := ""
		for ii := 0; ii < 8; ii++ {
			if cidrInt > 0 {
				tmp = tmp + "1"
				cidrInt--
			} else {
				tmp = tmp + "0"
			}
		}
		_int, err := strconv.ParseUint(tmp, 2, 64)
		if err != nil {
			return "", err
		}
		maskList = append(maskList, strconv.FormatUint(_int, 10))
	}
	netMask = strings.Join(maskList, ".")
	return netMask, nil
}

package ip_helpers

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type IPv4 [4]int

// ToIPv4 converts a string to a IPv4.
func ToIPv4(ip string) IPv4 {
	var newIP IPv4
	ipS := strings.Split(ip, ".")
	for i, v := range ipS {
		newIP[i], _ = strconv.Atoi(v)
	}
	return newIP
}

// ToString converts an IP from IPv4 type to string.
func (ip *IPv4) ToString() string {
	ipStringed := strconv.Itoa(ip[0])
	for i := 1; i < 4; i++ {
		strI := strconv.Itoa(ip[i])
		ipStringed += "." + strI
	}
	return ipStringed
}

// IsValid checks an IP address as valid or not.
func (ip *IPv4) IsValid() bool {
	for i, oct := range ip {
		if i == 0 || i == 3 {
			if oct < 1 || oct > 254 {
				return false
			}
		} else {
			if oct < 0 || oct > 255 {
				return false
			}
		}
	}
	return true
}

// PlusPlus increments an IPv4 value.
func (ip *IPv4) PlusPlus() *IPv4 {
	if ip[3] < 254 {
		ip[3] = ip[3] + 1
	} else {
		if ip[2] < 255 {
			ip[2] = ip[2] + 1
			ip[3] = 1
		} else {
			if ip[1] < 255 {
				ip[1] = ip[1] + 1
				ip[2] = 1
				ip[3] = 1
			} else {
				if ip[0] < 255 {
					ip[0] = ip[0] + 1
					ip[1] = 1
					ip[2] = 1
					ip[3] = 1
				}
			}
		}
	}
	return ip
}

// ParseIPSequence gets a sequence of IP addresses correspondent from an
// "init-end" entry.
func ParseIPSequence(ipSequence string) []IPv4 {
	var arrayIps []IPv4
	series, _ := regexp.Compile("([0-9]+)")
	// For sequence ips, using '-'
	lSeries := series.FindAllStringSubmatch(ipSequence, -1)
	for i := types.ToInt(lSeries[3][0]); i <= types.ToInt(lSeries[4][0]); i++ {
		arrayIps = append(arrayIps, IPv4{
			types.ToInt(lSeries[0][0]),
			types.ToInt(lSeries[1][0]),
			types.ToInt(lSeries[2][0]),
			i})
	}
	return arrayIps
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
func SubnetString(cidr int) (string, error) {
	cid := types.ToString(cidr)
	var maskList []string
	var netMask string
	cidrInt, err := strconv.ParseUint(cid, 10, 8)
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
		n, err := strconv.ParseUint(tmp, 2, 64)
		if err != nil {
			return "", err
		}
		maskList = append(maskList, strconv.FormatUint(n, 10))
	}
	netMask = strings.Join(maskList, ".")
	return netMask, nil
}

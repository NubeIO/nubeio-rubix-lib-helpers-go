package networking

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

/*
Check internet connections
*/

type Check struct {
	IpAddress string `json:"ip_address,omitempty"`
	Message   string `json:"message,omitempty"`
	Ok        bool   `json:"ok"`
}

//CheckInternetByInterface check internet connection for a port (will ping google.com)
func (nets *nets) CheckInternetByInterface(iface string) (connection Check, err error) {
	cmd := fmt.Sprintf("if ping -I %s -c 2 google.com; then echo OK; else echo DEAD ;fi", iface)
	ping, err := command.Run("bash", "-c", cmd)
	if err != nil {
		return
	}
	connection.Message = "service not known"
	if strings.Contains(ping, "OK") {
		connection.Message = "ok"
		connection.Ok = true
	} else if strings.Contains(ping, "unknown iface") {
		connection.Message = "failed to find network interface"
	} else if strings.Contains(ping, "Name or service not known") {
		connection.Message = "service not known"
	}
	return
}

// GetInternetIP fetches external IP address in ipv4 format
func (nets *nets) GetInternetIP() (connection Check, err error) {
	resp, err := http.Get("https://checkip.amazonaws.com")
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	content, _ := ioutil.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return
	}
	connection.IpAddress = strings.TrimRight(string(content), "\r\n")
	connection.Ok = true
	return
}

type NetInterface struct {
	Name         string   // Network interface name
	MTU          int      // MTU
	HardwareAddr string   // Hardware address
	Addresses    []string // Array with the network interface addresses
	Subnets      []string // Array with CIDR addresses of this network interface
	Flags        string   // Network interface flags (up, broadcast, etc)
}

func (nets *nets) GetValidNetInterfacesForWeb() ([]NetInterface, error) {
	ifaces, err := nets.GetValidNetInterfaces()
	if err != nil {
		return nil, errors.New("couldn't get interfaces")
	}
	if len(ifaces) == 0 {
		return nil, errors.New("couldn't find any legible interface")
	}
	var netInterfaces []NetInterface
	for _, iface := range ifaces {
		addrs, e := iface.Addrs()
		if e != nil {
			return nil, errors.New("failed to get addresses for interface")
		}
		netIface := NetInterface{
			Name:         iface.Name,
			MTU:          iface.MTU,
			HardwareAddr: iface.HardwareAddr.String(),
		}
		if iface.Flags != 0 {
			netIface.Flags = iface.Flags.String()
		}
		// Collect network interface addresses
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				// not an IPNet, should not happen
				return nil, fmt.Errorf("got iface.Addrs() element %s that is not net.IPNet, it is %T", addr, addr)
			}
			// ignore link-local
			if ipNet.IP.IsLinkLocalUnicast() {
				continue
			}
			netIface.Addresses = append(netIface.Addresses, ipNet.IP.String())
			netIface.Subnets = append(netIface.Subnets, ipNet.String())
		}
		// Discard interfaces with no addresses
		if len(netIface.Addresses) != 0 {
			netInterfaces = append(netInterfaces, netIface)
		}
	}
	return netInterfaces, nil
}

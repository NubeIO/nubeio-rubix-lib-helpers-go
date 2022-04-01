package edgeip

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/validation"
	"strings"
)

type EdgeNetworking struct {
	IPAddress  string `json:"ip_address" post:"true"`
	SubnetMask string `json:"subnet_mask" post:"true"`
	Gateway    string `json:"gateway" post:"true"`
	SetDHCP    bool   `json:"set_dhcp" post:"true"`
	RunAsSudo  bool   `json:"run_as_sudo"`
	Password   string `json:"-"`
}

func getInterface() (ok bool, interfaceName string, err error) {
	out, err := command.Run("connmanctl services")
	if err != nil {
		return false, "", errors.New("failed to get interface")
	} else {
		if strings.Contains(out, "*AO") {
			res := strings.ReplaceAll(out, "*AO Wired", "")
			return true, str.StandardizeSpaces(res), nil
		} else {
			return false, "", errors.New("failed to parse interface")
		}
	}
}

func SetIP(net EdgeNetworking) (ok bool, err error) {
	//const setIP = `sudo connmanctl config ${iface} --ipv4 manual ${ipAddress} ${subnetMask} ${gateway} --nameservers 8.8.8.8 8.8.4.4`;
	//const setIpDHCP = `sudo connmanctl config ${iface} --ipv4 dhcp`;

	if !net.SetDHCP {
		_, err = validation.IsIPAddr(net.IPAddress)
		if err != nil {
			return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an IPAddress", net.IPAddress))
		}
		_, err = validation.IsIPAddr(net.SubnetMask)
		if err != nil {
			return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an SubnetMask", net.SubnetMask))
		}
		_, err = validation.IsIPAddr(net.Gateway)
		if err != nil {
			return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an Gateway", net.Gateway))
		}
	}

	ok, iface, err := getInterface()
	if err != nil {
		return false, err
	}
	var setIpDHCP []string
	if net.RunAsSudo {
		setIpDHCP = []string{
			"sudo connmanctl config", iface, "--ipv4 dhcp",
		}
		if !net.SetDHCP {
			setIpDHCP = []string{
				"connmanctl config", iface, "--ipv4 manual", iface, net.IPAddress, net.SubnetMask, net.Gateway,
			}
		}
	} else {
		setIpDHCP = []string{
			"sudo connmanctl config", iface, "--ipv4 dhcp",
		}
		if !net.SetDHCP {
			setIpDHCP = []string{
				"connmanctl config", iface, "--ipv4 manual", iface, net.IPAddress, net.SubnetMask, net.Gateway,
			}
		}
	}

	_, err = command.Run(setIpDHCP...)
	if err != nil {
		return false, errors.New(" failed to update ip address")
	} else {
		return true, nil
	}
}

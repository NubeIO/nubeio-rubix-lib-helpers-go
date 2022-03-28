package networking

import (
	log "github.com/sirupsen/logrus"
)

//GetSubnet returns 255.255.255.0
func (nets *nets) GetSubnet(iFaceName string) (subnet string, err error) {
	net, err := nets.GetNetworkByIface(iFaceName)
	if err != nil {
		log.Errorf("Could not get network interfaces info: %v", err)
		return
	}
	subnet = net.NetMask
	return
}

//GetSubnetCIDR returns 24
func (nets *nets) GetSubnetCIDR(iFaceName string) (cidr int, err error) {
	net, err := nets.GetNetworkByIface(iFaceName)
	if err != nil {
		log.Errorf("Could not get network interfaces info: %v", err)
		return
	}
	cidr = net.NetMaskLength
	return
}

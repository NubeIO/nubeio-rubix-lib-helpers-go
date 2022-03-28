package networking

import (
	"errors"
	"net"
)

/*
Interfaces
*/

type InterfaceNames struct {
	Names []string `json:"interface_names"`
}

func (nets *nets) GetValidNetInterfaces() (interfaces []net.Interface, err error) {
	iFaces, err := net.Interfaces()
	for i := range iFaces {
		interfaces = append(interfaces, iFaces[i])
	}
	return
}

func (nets *nets) GetInterfacesNames() (interfaces InterfaceNames, err error) {
	i, err := nets.GetValidNetInterfaces()
	if err != nil {
		return interfaces, errors.New("couldn't get interfaces")
	}
	for _, n := range i {
		interfaces.Names = append(interfaces.Names, n.Name)
	}
	return
}

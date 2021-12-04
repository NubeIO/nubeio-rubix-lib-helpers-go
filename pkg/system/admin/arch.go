package admin

import (
	"strings"
)

type Arch struct {
	IsBeagleBone bool `json:"is_beagle_bone"`
	IsRaspberry  bool `json:"is_raspberry"`
	IsArm        bool `json:"is_arm"`
	IsAMD64      bool `json:"is_amd64"`
	IsAMD32      bool `json:"is_amd32"`
}

//DetectArch can detect hardware type is in ARM or AMD and also if hardware is for example a Raspberry PI
func (a *Admin) DetectArch() (arch Arch, ok bool, err error) {
	cmd := "tr '\\0' '\\n' </proc/device-tree/model;arch &&  dpkg --print-architecture"
	a.Host.CommandOpts.CMD = cmd
	o, ok, err := a.Host.RunCommand()
	if strings.Contains(o, "Raspberry Pi") {
		arch.IsRaspberry = true
		arch.IsArm = true
		return arch, true, err
	} else if strings.Contains(o, "BeagleBone Black") {
		arch.IsBeagleBone = true
		arch.IsArm = true
		return arch, true, err
	} else if strings.Contains(o, "amd64") {
		arch.IsAMD64 = true
		return arch, true, err
	} else if strings.Contains(o, "amd32") {
		arch.IsAMD32 = true
		return arch, true, err
	}
	return arch, false, err
}

package remote

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"strings"
)

type Arch struct {
	ArchModel    string `json:"arch_model"`
	IsBeagleBone bool   `json:"is_beagle_bone,omitempty"`
	IsRaspberry  bool   `json:"is_raspberry,omitempty"`
	IsArm        bool   `json:"is_arm,omitempty"`
	IsAMD64      bool   `json:"is_amd64,omitempty"`
	IsAMD32      bool   `json:"is_amd32,omitempty"`
	IsARMf       bool   `json:"is_armf,omitempty"`
	IsArmv7l     bool   `json:"is_armv7l,omitempty"`
}

//DetectArch can detect hardware type is in ARM or AMD and also if hardware is for example a Raspberry PI
func (a *Admin) DetectArch() (arch Arch, ok bool, err error) {
	cmd := "tr '\\0' '\\n' </proc/device-tree/model;arch &&  dpkg --print-architecture"
	a.Host.CommandOpts.CMD = cmd
	o, ok, err := a.Host.RunCommand()
	if err != nil || strings.Contains(o, " No such file or directory") {
		cmd = "dpkg --print-architecture"
		a.Host.CommandOpts.CMD = cmd
		o, ok, err = a.Host.RunCommand()
		arch.ArchModel = o
		if err != nil {
			return arch, false, err
		}
	}
	o = str.RemoveNewLine(o)
	if strings.Contains(o, "Raspberry Pi") {
		arch.IsRaspberry = true
		arch.IsArm = true
		return arch, true, err
	} else if strings.Contains(o, "BeagleBone Black") {
		arch.ArchModel = "BeagleBone Black"
		arch.IsBeagleBone = true
		arch.IsArm = true
		return arch, true, err
	} else if strings.Contains(o, "amd64") {
		arch.ArchModel = o
		arch.IsAMD64 = true
		return arch, true, err
	} else if strings.Contains(o, "amd32") {
		arch.ArchModel = o
		arch.IsAMD32 = true
		return arch, true, err
	} else if strings.Contains(o, "armhf") {
		arch.ArchModel = o
		arch.IsARMf = true
		arch.IsArm = true
		return arch, true, err
	} else if strings.Contains(o, "armv7l") {
		arch.ArchModel = o
		arch.IsArmv7l = true
		arch.IsArm = true
		return arch, true, err
	}
	return arch, false, err
}

package admin

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Node struct {
	IsInstalled      bool   `json:"is_installed"`
	InstalledVersion string `json:"installed_version"`
}

func (a *Admin) NodeGetVersion() (Node Node, ok bool, err error) {
	cmd := "nodejs -v"
	a.Host.CommandOpts.CMD = cmd
	o, ok, err := a.Host.RunCommand()
	if strings.Contains(o, "v") {
		Node.InstalledVersion = str.RemoveNewLine(o)
		Node.IsInstalled = true
		return Node, true, err
	} else {
		Node.InstalledVersion = o
		Node.IsInstalled = false
		return Node, true, err
	}
}

type NodeJSInstall struct {
	AlreadyInstalled bool   `json:"already_installed"`
	InstalledOk      bool   `json:"installed_ok"`
	TextOut          string `json:"text_out"`
}

func (a *Admin) InstallNode14() (NodeJSInstall NodeJSInstall, err error) {
	cmd := "sudo apt update -y \\\n  && curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash - \\\n  && sudo apt-get install -y nodejs \\\n  && nodejs -v"
	a.Host.CommandOpts.CMD = cmd
	out, ok, err := a.Host.RunCommand()
	if err != nil {
		log.Error("ufw: Install Error: ", err)
		NodeJSInstall.TextOut = out
		NodeJSInstall.InstalledOk = ok
		return NodeJSInstall, err
	}
	_, _, err = a.NodeGetVersion()
	if err != nil {
		log.Error("node: NodeGetVersion Error: ", err)
		NodeJSInstall.TextOut = ""
		NodeJSInstall.InstalledOk = false
		return NodeJSInstall, err
	}
	NodeJSInstall.AlreadyInstalled = true
	return NodeJSInstall, err

}

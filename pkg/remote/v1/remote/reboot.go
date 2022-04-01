package remote

import (
	log "github.com/sirupsen/logrus"
)

func (a *Admin) HostReboot() (ok bool, err error) {
	cmd := "sudo shutdown -r now"
	a.Host.CommandOpts.CMD = cmd
	_, ok, err = a.Host.RunCommand()
	if err != nil {
		log.Error("admin: HostReboot Error: ", err)
		return ok, err
	}
	return ok, err
}

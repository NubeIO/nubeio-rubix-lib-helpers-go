package ssh

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type Host struct {
	ID                   string    `json:"id" gorm:"primarykey"`
	Name                 string    `json:"name"  gorm:"type:varchar(255);unique;not null"`
	IP                   string    `json:"ip"`
	Port                 int       `json:"port"`
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	RubixPort            int       `json:"rubix_port"`
	RubixUsername        string    `json:"rubix_username"`
	RubixPassword        string    `json:"rubix_password"`
	RubixToken           string    `json:"-"`
	RubixTokenLastUpdate time.Time `json:"-"`
	IsLocalhost          bool      `json:"is_localhost"`
	SSH                  *goph.Client
	CommandOpts          CommandOpts
}

type CommandOpts struct {
	CMD         string
	Sudo, Debug bool
}

type serverSettings struct {
	Addr, Key, User, Catalog, Password string
	Port                               uint
}

type Controller struct {
	SSH *goph.Client
}

func (h Host) RunCommand() (out string, result bool, err error) {
	if h.IsLocalhost {
		out, err := command.RunCMD(h.CommandOpts.CMD, h.CommandOpts.Debug)
		if err != nil {
			return "", false, err
		}
		return string(out), false, err
	} else {
		c, err := h.newRemoteClient(h)
		if err != nil {
			return "", false, err
		}
		defer c.Close()
		o, err := c.Run(h.CommandOpts.CMD)
		if err != nil {
			return "", false, err
		}
		return string(o), true, err

	}
}

func (h *Host) newRemoteClient(host Host) (c *goph.Client, err error) {
	var cli serverSettings
	cli.Addr = host.IP
	cli.User = host.Username
	cli.Password = host.Password
	cli.Port = uint(host.Port)
	c, err = goph.NewConn(&goph.Config{
		User:     cli.User,
		Addr:     cli.Addr,
		Port:     cli.Port,
		Auth:     goph.Password(cli.Password),
		Callback: verifyHost,
	})
	return c, err
}

func verifyHost(host string, remote net.Addr, key ssh.PublicKey) error {
	// hostFound: is host in known hosts file.
	// err: error if key not in known hosts file OR host in known hosts file but key changed!
	hostFound, err := goph.CheckKnownHost(host, remote, key, "")
	// Host in known hosts but key mismatch!
	// Maybe because of MAN IN THE MIDDLE ATTACK!
	if hostFound && err != nil {

		return err
	}

	// handshake because public key already exists.
	if hostFound && err == nil {

		return nil
	}
	return goph.AddKnownHost(host, remote, key, "")
}

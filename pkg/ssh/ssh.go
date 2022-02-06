package ssh

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	sh "github.com/helloyi/go-sshclient"
	"github.com/melbahja/goph"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type Host struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
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

//fmt.Println("222222")
//client1, err := sh.DialWithPasswd("120.151.62.75:2221", "debian", "N00BConnect")
//if err != nil {
//fmt.Println(err)
//}
//defer client1.Close()
//
//ccc, _ := client1.Cmd("pwd").Output()
//fmt.Println(string(ccc))

//RunCommand will run a local or remote command, if CommandOpts.Sudo is true then a sudo is added to the existing command (cmd = "sudo " + CommandOpts.CMD)
func (h Host) RunCommand() (out string, result bool, err error) {
	cmd := h.CommandOpts.CMD
	if h.CommandOpts.Sudo {
		cmd = "sudo " + h.CommandOpts.CMD
	}
	if h.IsLocalhost {
		out, err := command.RunCMD(cmd, h.CommandOpts.Debug)
		if err != nil {
			return "", false, err
		}
		return string(out), false, err
	} else {
		host := fmt.Sprintf("%s:%d", h.IP, h.Port)
		c, err := sh.DialWithPasswd(host, h.Username, h.Password)
		if err != nil {
			return "", false, err
		}
		defer c.Close()
		o, err := c.Cmd(cmd).Output()
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
	log.Println("SSH", hostFound, err)
	return goph.AddKnownHost(host, remote, key, "")
}

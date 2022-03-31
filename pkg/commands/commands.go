package commands

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/ssh"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

/*Commands
Parameters:
	: Command
	: RemoteCommand
	: VerifyHost
	: Host  ssh.Host
Description:
*/
type Commands struct {
	cmd           *exec.Cmd
	Command       string
	Args          []string
	Shell         bool
	RemoteCommand bool
	VerifyHost    bool
	Host          ssh.Host
	Debug         bool
}

type Result struct {
	resultByte []byte
	result     string
	err        error
}

func New(e *Commands) *Commands {
	return e
}

func (inst *Commands) RunCommand() *Result {
	return inst.runCommand()
}

func (inst *Commands) ChainCommand(cmd string) *Commands {
	inst.Command = cmd
	inst.runCommand()
	return inst
}

func (inst *Commands) Result() *Result {
	return res
}

func (res *Result) Error() error {
	return res.err
}

func (res *Result) AsString() string {
	return res.result
}

func debug(debug bool) {
	if !debug {
		return
	}
	if res.err != nil {
		log.Errorln("----------nube-helpers-commands-error----------")
		log.Errorln(res.err)
	} else {
		log.Println("----------nube-helpers-commands----------")
		log.Println(res.result)
	}
}

func (res *Result) Log() {
	debug(true)
}

func (res *Result) AsByte() []byte {
	return res.resultByte
}

var res = &Result{}

func (inst *Commands) runCommand() *Result {
	cmd := inst.Command
	if !inst.RemoteCommand {
		out, err := cmdRun(cmd, false)
		res.resultByte = out
		res.result = string(out)
		res.err = err
		if err != nil {
			debug(inst.Debug)
			return res
		}
		debug(inst.Debug)
		return res
	} else {
		//host := fmt.Sprintf("%s:%d", inst.Host.IP, inst.Host.Port)
		//c, err := sh.DialWithPasswd(host, inst.Host.Username, inst.Host.Password)
		//if err != nil {
		//	return "", false, err
		//}
		//defer c.Close()
		//o, err := c.Cmd(cmd).Output()
		//if err != nil {
		//	return "", false, err
		//}
		//return string(o), true, err
	}
	return res
}

//CommandsList are pre-made commands to be used with New(&Commands{Command: CommandsList.GetHomeDir}).RunCommand().AsString()
var CommandsList = struct {
	GetHomeDir, PWD string
}{
	GetHomeDir: "echo $HOME",
	PWD:        "pwd",
}

func cmdRun(sh string, debug bool) ([]byte, error) {
	cmd := exec.Command("bash", "-c", sh)
	res, e := cmd.CombinedOutput()

	if debug {
		fmt.Printf("[admin debug] %s\n", cmd.String())
	}
	if e != nil {
		defer cmd.Process.Kill()
		return nil, e
	}

	defer cmd.Process.Kill()
	return res, e
}

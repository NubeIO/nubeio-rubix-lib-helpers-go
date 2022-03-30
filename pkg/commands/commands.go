package commands

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/ssh"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"os/exec"
)

type Commands struct {
	cmd           *exec.Cmd
	Command       string
	Args          []string
	Shell         bool
	RemoteCommand bool
	CheckHost     bool
	Host          ssh.Host
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

func (res *Result) Error() error {
	return res.err
}

func (res *Result) AsString() string {
	return res.result
}

func (res *Result) AsByte() []byte {
	return res.resultByte
}

func (inst *Commands) runCommand() *Result {
	cmd := inst.Command
	res := &Result{}
	if !inst.RemoteCommand {
		out, err := command.RunCMD(cmd, false)
		res.resultByte = out
		res.result = string(out)
		res.err = err
		if err != nil {
			return res
		}
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

func (inst *Commands) GetHomeDir() *Result {
	inst.Command = "echo $HOME"
	return inst.runCommand()
}

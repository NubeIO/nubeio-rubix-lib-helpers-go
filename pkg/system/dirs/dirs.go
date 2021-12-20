package dirs

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/ssh"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

//Dirs for local or remote dir functions
type Dirs struct {
	Name          string   `json:"name"`
	CheckIfExists bool     `json:"check_if_exists"`
	IfExistsClear bool     `json:"if_exists_clear"`
	Host          ssh.Host `json:"host"`
}

//MKDir make a new dir, with option to check if existing and if so then also to clear its contents
func (d *Dirs) MKDir() (out string, err error) {
	dirExists := false
	if d.CheckIfExists {
		check, err := d.DirExists()
		if err != nil {
			return "", err
		}
		dirExists = check
		if check && d.IfExistsClear {
			_, err := d.WipeDir()
			if err != nil {
				return "", err
			}
		}
	}
	if !dirExists {
		d.Host.CommandOpts.CMD = fmt.Sprintf("mkdir %s", d.Name)
		check, _, err := d.Host.RunCommand()
		if err != nil {
			return "", err
		}
		return check, nil
	}
	return "", nil
}

//DirExists check if a dir exists
func (d *Dirs) DirExists() (out bool, err error) {
	d.Host.CommandOpts.CMD = fmt.Sprintf("if [ -d %s ]; then echo 'yes'; else echo 'no'; fi", d.Name)
	dirExists := false
	check, _, err := d.Host.RunCommand()
	if err != nil {
		return false, err
	}
	if strings.Contains(check, "yes") {
		dirExists = true
	}
	log.Println("dirExists()", "result: ", check)
	return dirExists, nil
}

//WipeDir is to remove all file's and dirs inside the dir
func (d *Dirs) WipeDir() (out bool, err error) {
	d.Host.CommandOpts.CMD = fmt.Sprintf("rm -r %s/*", d.Name)
	_, _, err = d.Host.RunCommand()
	if err != nil {
		log.Println("wipeDir()", "result:", "fail", "command:", d.Host.CommandOpts.CMD)
		return false, nil
	}
	log.Println("wipeDir()", "result:", "pass")
	return true, nil

}

func DirRemoveContents(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func DirIsWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

//DirChangePermissions
/*
@param path /etc
@param permissions 0700
*/
func DirChangePermissions(path string, permissions uint32) (ok bool, err error) {
	err = os.Chmod(path, os.FileMode(permissions))
	if err != nil {
		log.Error(err)
		return false, err
	}
	return true, err
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func GetUserHomeDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, err
}

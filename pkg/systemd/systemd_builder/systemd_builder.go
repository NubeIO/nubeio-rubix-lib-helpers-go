package systemd_builder

import (
	"fmt"
	"io/ioutil"
)

func build() {

	name := "aidans-service"
	user := "aidan"
	directory := "/home/aidan"
	execCmd := "/usr/bin/python3 something.py"

	template := `[Unit]
Description=%v Service
After=network.target
[Service]
User=%v
WorkingDirectory=%v
ExecStart=%v
Restart=always
[Install]
WantedBy=multi-user.target`

	serviceFile := fmt.Sprintf(template, name, user, directory, execCmd)

	servicePath := fmt.Sprintf("/home/aidan/%v.service", name)
	ioutil.WriteFile(servicePath, []byte(serviceFile), 0644)

	//exec.Command("systemctl", "enable", name).Run()
	//exec.Command("service", name, "restart").Run()

	fmt.Println("Service File Created.")
}

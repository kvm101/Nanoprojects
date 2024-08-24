package hcmd

import (
	"os/exec"
	"log"
	"strings"
)

func Exec(command string) {
	parts_cmd := strings.Fields(command)
	cmd := exec.Command(parts_cmd[0], parts_cmd[1:]...)
	
	cmd_err := cmd.Run()
	if cmd_err != nil {
		log.Println(cmd_err)
	}
}
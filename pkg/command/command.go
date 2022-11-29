package command

import "os/exec"

func Exec(command string) error {
	cmd := exec.Command(command)

	return cmd.Run()
}

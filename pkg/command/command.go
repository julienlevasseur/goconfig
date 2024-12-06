package command

import (
	"fmt"
	"os/exec"
)

// func Exec(command string, args ...string) error {
func Exec(command string, args []string) error {
	cmd := exec.Command(command, args...)

	return cmd.Run()
}

func ExecWithOutput(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(cmd.Stdout)
		return "", err
	}

	return string(out), nil
}

package command

import "os/exec"

// func Exec(command string, args ...string) error {
func Exec(command string) error {
	//var a []string
	//for _, i := range args {
	//	a = append(a, i)
	//}
	cmd := exec.Command(command)

	return cmd.Run()
}

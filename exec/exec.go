package exec

import "os/exec"

func Execute(name string, args ...string) (err error) {
	c := exec.Command(name, args...)
	err = c.Run()
	return err
}

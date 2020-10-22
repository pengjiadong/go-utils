package exec

import (
	"os/exec"

	"github.com/pengjiadong/go-utils/str"
	"github.com/sirupsen/logrus"
)

func Execute(name string, arg ...string) (content string, err error) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		logrus.Errorf("DoCommand error: %s;%s\n", err, arg)
	}
	content = str.GBK2UTF8(output)
	return
}

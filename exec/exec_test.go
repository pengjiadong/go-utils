package exec

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	x, y := Execute("ping", "www.baidu.com")
	fmt.Println(x)
	fmt.Println(y)
}

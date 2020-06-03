package golang

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	file := "/Users/patrick/GoProjects/dubbo-samples/golang/helloworld/dubbo/docker/docker-compose.yaml"
	cmd := exec.Command("docker-compose", "-f", file, "up", "--build", "--abort-on-container-exit")
	var out bytes.Buffer
	cmd.Stdout = &out
	println(cmd.String())
	err := cmd.Run()
	assert.Nil(t, err)
	str := out.String()
	println(str)
	i := strings.Index(str, "response result: &{A001 Alex Stocks 18 2013-01-02 00:00:00 +0000 UTC}")
	assert.NotEqual(t, i, -1)
}

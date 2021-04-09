package imp

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func ExecCommand(commandString string) bool {
	cmd := exec.Command("/usr/bin/bash", "-c", commandString)
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		fmt.Println(line)
	}
	cmd.Wait()
	return true
}

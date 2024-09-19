package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func main() {
	newScaner := bufio.NewScanner(os.Stdin)
	args := strings.Join(os.Args[2:], " ")
	for newScaner.Scan() {
		line := newScaner.Text()
		cmd := exec.Command(os.Args[1], args, line)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

}

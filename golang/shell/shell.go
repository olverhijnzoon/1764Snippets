package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Shell")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">1764Shell>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		executeCommand(input)
	}
}

func executeCommand(command string) {
	args := strings.Fields(command)
	// essentially a wrapper around this
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A command is required")
		os.Exit(1)
	}

	i := 1
	command := os.Args[i]
	flags := []string{}

	rawTime := false
	stdOut := true

	for command[0] == '-' {
		if command == "-r" {
			rawTime = true
		} else if command == "-ns" {
			stdOut = false
		}

		i += 1
		command = os.Args[i]
	}

	if len(os.Args) > 1+i {
		flags = os.Args[1+i:]
	}

	cmd := exec.Command(command, flags...)
	cmd.Stdin = os.Stdin

	if stdOut {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	cmd.Env = os.Environ()

	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)

	if err != nil {
		fmt.Println(err)
	}

	ms := duration.Milliseconds()

	if rawTime {
		fmt.Println(duration.Microseconds())
	} else {
		if stdOut {
			fmt.Println()
		}

		fmt.Println("Execution time:")

		if ms > 0 {
			fmt.Printf("  %dms\n", ms)
		} else {
			fmt.Printf("  0.%dms (%dns)\n", duration.Microseconds(), duration.Nanoseconds())
		}
	}
}

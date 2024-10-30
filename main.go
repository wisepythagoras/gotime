package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const NAME = "gotime"
const VERSION = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Run \"%s -h\" to see how to use this command\n", NAME)
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
		} else if command == "-h" {
			fmt.Printf("Usage: %s [OPTION]... COMMAND [ARGS]...\n", NAME)
			fmt.Printf("Examples:\n  %s tar -cavf target.tar.xz directory\n", NAME)
			fmt.Printf("  cat /var/log/auth.log | %s grep 'pattern'\n", NAME)
			fmt.Println()
			fmt.Println("Measure the execution time of a program")
			fmt.Println()
			fmt.Println("Available flags")
			fmt.Println("  -r\tDisplay only the raw microseconds measurement")
			fmt.Println("  -ns\tPrevent any output from reaching stdout or stderr")
			fmt.Println("  -v\tShow the version of this program")
			os.Exit(0)
		} else if command == "-v" {
			fmt.Printf("%s %s\n", NAME, VERSION)
			os.Exit(0)
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

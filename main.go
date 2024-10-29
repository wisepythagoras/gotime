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

	command := os.Args[1]
	flags := []string{}

	if len(os.Args) > 2 {
		flags = os.Args[2:]
	}

	cmd := exec.Command(command, flags...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)

	if err != nil {
		fmt.Println(err)
	}

	ms := duration.Milliseconds()

	fmt.Println()
	fmt.Println("Execution time:")

	if ms > 0 {
		fmt.Printf("  %dms\n", ms)
	} else {
		fmt.Printf("  0.%dms (%dns)\n", duration.Microseconds(), duration.Nanoseconds())
	}
}

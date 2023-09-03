package main

import (
	"os"
)

func main() {
	cli := &CLI{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	exitcode := cli.Run(os.Args)
	os.Exit(exitcode)
}

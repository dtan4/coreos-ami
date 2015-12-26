package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		channel string
		region  string
		amiType string

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&channel, "c", "", "CoreOS channel (alpha, beta, stable)")
	flags.StringVar(&region, "r", "", "AWS region")
	flags.StringVar(&amiType, "t", "", "AMI type (PV, HVM)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = region

	_ = amiType

	if channel == "" {
		fmt.Fprintln(cli.errStream, "CoreOS channel (-c) must be specified!")
		fmt.Fprintln(cli.errStream, "Available channels: alpha, beta, stable")
		return ExitCodeError
	}

	return ExitCodeOK
}

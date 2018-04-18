package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/devonartis/s3cli/command"
)

// ExitErrorf prints out error using a custom string message
func ExitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "profile",
		Value: "default",
		Usage: "aws profile to use",
	},
	cli.StringFlag{
		Name:  "region",
		Value: "us-east-1",
		Usage: "aws region to use",
	},
}

var Commands = []cli.Command{
	{
		Name:   "listbuckets",
		Usage:  "List all buckets",
		Action: command.CmdListbuckets,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "print",
				Usage: "print list of buckets to csv file",
			},
		},
	},
	{
		Name:   "deletebucket",
		Usage:  "delete bucket",
		Action: command.CmdDeletebucket,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound function will trigger if Command is not found
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

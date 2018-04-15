package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/devonartis/s3cli/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "listbuckets",
		Usage:  "List all buckets",
		Action: command.CmdListbuckets,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "deletebucket",
		Usage:  "delete bucket",
		Action: command.CmdDeletebucket,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

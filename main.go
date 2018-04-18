package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	const Name string = "s3cli"
	const Version string = "0.1.0"
	app := cli.NewApp()

	app.Name = Name
	app.Version = Version
	app.Author = "T. Devon Artis"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)

}

package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/foostan/bravia-go/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_MITCHELH_CLI",
		Name:   "mitchelh_cli",
		Value:  "",
		Usage:  "",
	},
}

var Commands = []cli.Command{
	{
		Name:   "auth",
		Usage:  "",
		Action: command.CmdAuth,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

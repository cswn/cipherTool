package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

type SubCommand interface {
	Name() string
	Flags(*flag.FlagSet)
	Description() string
	Run([]string)
}

var subcommands []SubCommand

func init() {
	subcommands = append(subcommands, &CeasarSubCommand{})
	subcommands = append(subcommands, &VigenereSubCommand{})
}

func usageForSubCommand(subcommand SubCommand) string {
	usage := "  - " + (color.InBlue(subcommand.Name()))
	usage += fmt.Sprintf("\n      %s\n", subcommand.Description())
	return usage
}

func usage() string {
	usage := "goCipher is a command line tool for encoding and decoding fun secret messages.\n"
	usage += "Available commands:\n"
	for _, subcommand := range subcommands {
		usage += usageForSubCommand(subcommand)
	}
	return usage
}

func Main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, color.InRed("You must provide a valid subcommand.\n"))
		fmt.Fprintf(os.Stderr, "%s\n", usage())
		return
	}

	subcommand := args[1]

	if subcommand == "help" {
		fmt.Fprintf(os.Stderr, "%s\n", usage())
		return
	}

	for _, cmd := range subcommands {
		if cmd.Name() == subcommand {
			flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
			cmd.Flags(flagSet)
			err := flagSet.Parse(args[2:])
			if err != nil {
				fmt.Fprintf(os.Stderr, color.InRed("There was an error: %s\n"), err.Error())
				os.Exit(1)
			}
			cmd.Run(args)
			return
		}
	}

	fmt.Fprintln(os.Stderr, color.InRed("You must provide a valid subcommand.\n"))
	fmt.Fprintf(os.Stderr, "%s\n", usage())
}

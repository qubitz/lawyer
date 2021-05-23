package commands

import (
	"errors"
	"flag"
)

type helpSetup struct {
	flags *flag.FlagSet
	print bool
}

func setupForHelp() setup {
	return new(helpSetup)
}

func (setup helpSetup) Parse(args []string) Command {
	switch len(args) {
	case 0:
		return new(helpCommand)
	case 1:
		return helpCommand{
			subject: args[0],
		}
	default:
		return suggestHelpWith(errors.New("too many arguments"))
	}
}

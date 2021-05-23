package commands

import (
	"errors"
	"flag"
)

type indictSetup struct {
	flags *flag.FlagSet
}

func setupForIndict() indictSetup {
	flags := flag.NewFlagSet("indict", flag.ExitOnError)

	return indictSetup{
		flags: flags,
	}
}

func (setup indictSetup) Parse(args []string) Command {
	setup.flags.Parse(args)

	paths := setup.flags.Args()
	if len(paths) == 0 {
		return suggestHelpWith(
			errors.New("defendant path(s) required for indictment"))
	}

	return &indictCommand{
		paths: paths,
	}
}

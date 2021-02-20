package commands

import (
	"errors"
	"fmt"
)

// A Command executes a predefined procedure established before its creation.
type Command interface {
	Execute() error
}

type setup interface {
	Parse(args []string) Command
}

// Parse creates a valid Command from args. The expected format
// of args is: "<subcommand> <arguments...>". If an error occurs and a
// valid command can not be constructed, a new command is created explaining
// the error and suggesting the "lawyer help" command.
func Parse(args []string) Command {
	if len(args) == 0 {
		return suggestHelp(errors.New("no subcommand specified"))
	}

	setup, err := getSetup(args[0])
	if err != nil {
		return suggestHelp(err)
	}

	return setup.Parse(args[1:])
}

func getSetup(command string) (setup, error) {
	switch command {
	case "indict":
		return setupForIndict(), nil
	case "help", "--help", "-h":
		return setupForHelp(), nil
	default:
		return nil, fmt.Errorf("unknown command \"%v\"", command)
	}
}

func suggestHelp(err error) suggestHelpCommand {
	return suggestHelpCommand{
		error: err,
	}
}

type suggestHelpCommand struct {
	error
}

func (suggHelp suggestHelpCommand) Execute() error {
	fmt.Println("invalid command: " + suggHelp.error.Error())
	fmt.Println("use \"lawyer help\" to show usage information")
	return nil
}

package commands

import (
	"fmt"

	"../constants"
)

// A Command executes a predefined procedure established before its creation.
type Command interface {
	Execute() error
}

type setup interface {
	Parse(args []string) Command
}

var helpSuggestion = "use \"lawyer help\" to show usage information"

// Parse creates a valid Command from args. The expected format
// of args is: "<subcommand> <arguments...>". If an error occurs and a
// valid command can not be constructed, a new command is created explaining
// the error and suggesting the "lawyer help" command.
func Parse(args []string) Command {
	if len(args) == 0 {
		return welcomeCommand{}
	}

	setup, err := getSetup(args[0])
	if err != nil {
		return suggestHelpWith(err)
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

func suggestHelpWith(err error) suggestHelpCommand {
	return suggestHelpCommand{
		error: err,
	}
}

type suggestHelpCommand struct {
	error
}

func (suggHelp suggestHelpCommand) Execute() error {
	fmt.Println("invalid command: " + suggHelp.error.Error())
	fmt.Println(helpSuggestion)
	return nil
}

type welcomeCommand struct{}

func (welcome welcomeCommand) Execute() error {
	fmt.Println(constants.LawyerDescription)
	fmt.Println("version " + constants.LawyerVersion)
	fmt.Println(helpSuggestion)
	return nil
}

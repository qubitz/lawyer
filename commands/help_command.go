package commands

import (
	"fmt"

	"github.com/qubitz/lawyer/commands/manuals"
	"github.com/qubitz/lawyer/constants"
)

type helpCommand struct {
	subject string
}

func (help helpCommand) Execute() (err error) {
	if help.subject == "" {
		printSummary()
	} else {
		err = printManual(help.subject)
	}

	return err
}

func printSummary() {
	fmt.Println(constants.LawyerDescription)
	fmt.Println()
	printUsage()
	fmt.Println()
	printCommands()
	fmt.Println()
	printTopics()
}

func printUsage() {
	fmt.Print("Usage:")
	fmt.Println()
	fmt.Println("\tlawyer <command> [arguments]")
}

func printCommands() {
	fmt.Println("The commands are:")
	fmt.Println()
	for command, manual := range manuals.CommandManualMap {
		fmt.Printf("\t%-10v   %v\n", command, manual.Description)
	}
	fmt.Println()
	fmt.Println("Use \"lawyer help <command>\" for more information about a command.")
}

func printTopics() {
	fmt.Println("Additional help topics:")
	fmt.Println()
	for topic, manual := range manuals.TopicManualMap {
		fmt.Printf("\t%-10v   %v\n", topic, manual.Description)
	}
	fmt.Println()
	fmt.Println("Use \"lawyer help <topic>\" for more information about that topic.")
}

func printManual(command string) error {
	foundManual := manuals.EntireManualMap[command]

	if foundManual == *new(manuals.Manual) {
		return fmt.Errorf("no manual found for \"%v\"", command)
	}

	fmt.Println(foundManual.Content)
	return nil
}

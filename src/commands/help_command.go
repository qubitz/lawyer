package commands

import (
	"fmt"

	"./manuals"
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
	fmt.Println("Lawyer is a tool for maintaining copyright headers in source code.")
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
		fmt.Printf("\t%-10v   %v\n", command, manual.Summary)
	}
	fmt.Println()
	fmt.Println("Use \"go help <command>\" for more information about a command.")
}

func printTopics() {
	fmt.Println("Additional help topics:")
	fmt.Println()
	for topic, manual := range manuals.TopicManualMap {
		fmt.Printf("\t%-10v   %v\n", topic, manual.Summary)
	}
	fmt.Println()
	fmt.Println("Use \"go help <topic>\" for more information about that topic.")
}

func printManual(command string) error {
	foundManual := manuals.EntireManualMap[command]

	if foundManual == *new(manuals.Manual) {
		return fmt.Errorf("no manual found for \"%v\"", command)
	}

	fmt.Println(foundManual.Content)
	return nil
}

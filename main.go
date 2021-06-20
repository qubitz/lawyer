package main

import (
	"fmt"
	"os"

	"github.com/TwinProduction/go-color"
	"github.com/qubitz/lawyer/commands"
)

func main() {
	err := commands.Parse(os.Args[1:]).Execute()

	if err != nil {
		fmt.Println(color.Ize(color.Red, "error while executing command"))
		fmt.Println(color.Ize(color.Red, err.Error()))
		os.Exit(1)
	}
}

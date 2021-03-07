package main

import (
	"fmt"
	"os"

	"./commands"
	"github.com/TwinProduction/go-color"
)

func main() {
	err := commands.Parse(os.Args[1:]).Execute()

	if err != nil {
		fmt.Println(color.Ize(color.Red, "error while executing command"))
		fmt.Println(color.Ize(color.Red, err.Error()))
		os.Exit(1)
	}
}

package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/qubitz/lawyer/laws"
	"github.com/qubitz/lawyer/trial"
	"github.com/TwinProduction/go-color"
)

type indictCommand struct {
	paths   []string
	lawPath string
}

func (indictment *indictCommand) Execute() error {
	law, err := laws.RetrieveFrom(indictment.lawPath)
	if err != nil {
		return fmt.Errorf("unable to retrieve law contents\n%w", err)
	}

	for _, path := range indictment.paths {
		err := indict(path, law.Expected)

		if err != nil {
			fmt.Print(color.Yellow)
			fmt.Printf("%v dismissed\n", path)
			fmt.Println(err.Error())
			fmt.Print(color.Reset)
		}
	}

	fmt.Println("indictment complete")
	return nil
}

func indict(path string, expected regexp.Regexp) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	innocent, evidence := trial.Conduct(expected, file)
	fmt.Printf("%v is %v\n", path, trial.ToVerdict(innocent))

	if !innocent {
		fmt.Println("\n" + trial.FormatEvidence(evidence))
	}

	return nil
}

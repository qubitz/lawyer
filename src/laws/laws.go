package laws

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// A Law describes the expected values of a file header in regular expressions.
type Law struct {
	Expected regexp.Regexp
}

// RetrieveFrom extracts a Law from filepath.
// If filepath is empty, it attempts to find a Law file at
// the current working directory or, if not found, returns an error.
// Otherwise, an error is returned if filepath can not be opened, is of
// the wrong file type, or contains invalid data. 
func RetrieveFrom(filepath string) (law Law, err error) {
	if filepath == "" {
		fmt.Print("searching for the law...")
		filepath, err = findFile()
		if err != nil {
			fmt.Print("\n")
			return law, err
		}
		fmt.Printf("found %v\n", filepath)
	}

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return law, err
	}

	return retrieveLaw(file)
}

func findFile() (string, error) {
	path, _ := os.Getwd()
	path = filepath.Join(path, "law.*")
	matches, _ := filepath.Glob(path)

	if len(matches) == 0 {
		return "", fmt.Errorf(
			`unable to find any "law" file in the current directory` + "\n" +
				`enter "lawyer help law" for more information`)
	}
	return matches[0], nil
}

func retrieveLaw(file *os.File) (Law, error) {
	ext := filepath.Ext(file.Name())
	switch ext {
	case ".yml", ".yaml":
		return retrieveFromYaml(file)
	default:
		return Law{}, fmt.Errorf(
			`law file format "%v" not supported`+"\n"+
				`enter "lawyer help law" for more information`, ext)
	}
}

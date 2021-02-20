package laws

import (
	"io"
	"regexp"

	"gopkg.in/yaml.v3"
)

type yamlLaw struct {
	Expected string
}

func retrieveFromYaml(reader io.Reader) (Law, error) {
	raw := new(yamlLaw)
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)
	
	err := decoder.Decode(raw)
	if err != nil {
		return Law{}, err
	}

	regex, err := regexp.Compile(raw.Expected)

	return Law{
		Expected: *regex,
	}, err
}

package trial

import "github.com/TwinProduction/go-color"

// ToVerdict converts isInnocent into a fancy string
// where isInnocent == true: green "innocent" and
// isInnocent == false: red "guilty".
func ToVerdict(isInnocent bool) string {
	if isInnocent {
		return color.Ize(color.Green, "innocent")
	}
	return color.Ize(color.Red, "guilty")
}

// FormatEvidence sandwiches evidence in a multi-line string.
func FormatEvidence(evidence string) (format string) {
	const lineSep = "--------------------------------------------------"
	format += color.Ize(color.Cyan, "evidence\n"+lineSep+"\n")
	format += color.Ize(color.Gray, evidence+"\n")
	format += color.Ize(color.Cyan, lineSep+"\n")
	return format
}

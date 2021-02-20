package trial

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// Conduct attempts to match the file header lines of reader and expected.
// Returns whether there was a valid match and the file header contents of
// reader.
func Conduct(expected regexp.Regexp, reader io.Reader) (bool, string) {
	evidence := getEvidence(reader)
	return expected.MatchString(evidence), evidence
}

func getEvidence(reader io.Reader) string {
	lines := getFileLines(reader, 10)
	return strings.Join(lines, "\n")
}

func getFileLines(reader io.Reader, lnCount int) []string {
	scanner := bufio.NewScanner(reader)

	lines := make([]string, lnCount)
	idx := 0
	for idx < lnCount && scanner.Scan() {
		lines[idx] = scanner.Text()
		idx++
	}

	return lines[:idx]
}

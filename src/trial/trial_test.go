package trial

import (
	"io"
	"regexp"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockReader struct {
	mock.Mock
}

func (reader *mockReader) Read(p []byte) (int, error) {
	returns := reader.Mock.Called(p)

	n := copy(p, returns.String(0))
	return n, returns.Error(1)
}

var cases = []struct {
	law      regexp.Regexp
	evidence string
	verdict  bool
}{
	{
		law:      compile("test"),
		evidence: "tset",
		verdict:  false,
	},
	{
		law:      compile("test"),
		evidence: "test",
		verdict:  true,
	},
	{
		law:      compile(`^test\dtest$`),
		evidence: "test0test",
		verdict:  true,
	},
}

func TestCorrectVerdicts(t *testing.T) {
	for _, c := range cases {
		got, evidence := Conduct(c.law, getReaderThatReads(c.evidence, nil))
		require.Equal(t, c.evidence, evidence, c)
		require.Equal(t, c.verdict, got, c)
	}
}

func compile(regex string) regexp.Regexp {
	r, e := regexp.Compile(regex)
	if e != nil {
		panic(e)
	}
	return *r
}

func getReaderThatReads(toRead string, err error) io.Reader {
	reader := new(mockReader)
	reader.Mock.On("Read", mock.AnythingOfType("[]uint8")).Return(toRead, io.EOF)
	return reader
}

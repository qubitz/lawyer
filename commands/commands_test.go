package commands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsesCommands(t *testing.T) {
	path := "tes.txt"

	var cases = []struct {
		in   []string
		want Command
	}{
		{
			in: []string{"indict", path},
			want: &indictCommand{
				paths:   []string{path},
				lawPath: "",
			},
		},
	}

	for _, cas := range cases {
		got := Parse(cas.in)
		require.Equal(t, cas.want, got)
	}
}

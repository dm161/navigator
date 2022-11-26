package command

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()
	for name, tc := range []struct {
		in  string
		out Command
		err error
	}{
		{
			in: "FFBB",
			out: Command{
				Sequence: []rune{'F', 'F', 'B', 'B'},
			},
		},
		{
			in:  "",
			err: ErrCommandCannotBeEmpty,
		},
		{
			in:  "H",
			err: ErrCommandInvalidValue,
		},
	} {
		t.Run(fmt.Sprintf("tc-%d", name), func(t *testing.T) {
			out, err := ParseCommand(tc.in)
			if err != tc.err {
				t.Errorf("ParseCommand expected error: %v, got: %v", tc.err, err)
			}
			if len(out.Sequence) != len(tc.out.Sequence) {
				t.Errorf("ParseCommand expected output: %v, got: %v", tc.out, out)
			}
		})
	}
}

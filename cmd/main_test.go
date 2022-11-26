package main

import (
	"bytes"
	"fmt"
	"testing"

	"diego.pizza/equalexperts/mars/internal/command"
	"diego.pizza/equalexperts/mars/internal/position"
)

func TestRun(t *testing.T) {
	t.Parallel()
	for name, tc := range []struct {
		start string
		cmd   string
		buf   *bytes.Buffer
		out   string
		err   error
	}{
		{
			start: "4,2,EAST",
			cmd:   "FLFFFRFLB",
			buf:   bytes.NewBuffer([]byte(``)),
			out:   "(6, 4, NORTH)",
		},
		{
			start: "4,",
			cmd:   "FLFFFRFLB",
			buf:   bytes.NewBuffer([]byte(``)),
			err:   position.ErrPositionInvalid,
		},
		{
			start: "4,2,B",
			cmd:   "FLFFFRFLB",
			buf:   bytes.NewBuffer([]byte(``)),
			err:   position.ErrInvalidDirectionValue,
		},
		{
			start: "4,2,EAST",
			cmd:   "",
			buf:   bytes.NewBuffer([]byte(``)),
			err:   command.ErrCommandCannotBeEmpty,
		},
		{
			start: "4,2,EAST",
			cmd:   "AA",
			buf:   bytes.NewBuffer([]byte(``)),
			err:   command.ErrCommandInvalidValue,
		},
	} {
		t.Run(fmt.Sprintf("tc-%d", name), func(t *testing.T) {
			err := run(tc.start, tc.cmd, tc.buf)
			if err != tc.err {
				t.Errorf("run expected error: %v, got: %v", tc.err, err)
			}
			if string(tc.buf.Bytes()) != tc.out {
				t.Errorf("run expected output: %v, got: %v", tc.out, string(tc.buf.Bytes()))
			}
		})
	}
}

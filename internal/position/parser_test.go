package position

import (
	"fmt"
	"testing"
)

func TestParsePosition(t *testing.T) {
	t.Parallel()
	for name, tc := range []struct {
		in  string
		out Position
		err error
	}{
		{
			in: "1,2,NORTH",
			out: Position{
				X:         1,
				Y:         2,
				Direction: DirectionNorth,
			},
		},
		{
			in:  "1,2,SUB",
			err: ErrInvalidDirectionValue,
		},
		{
			in:  "",
			err: ErrPositionCannotBeEmpty,
		},
		{
			in:  "1.2",
			err: ErrPositionInvalid,
		},
	} {
		t.Run(fmt.Sprintf("tc-%d", name), func(t *testing.T) {
			out, err := ParsePosition(tc.in)
			if err != tc.err {
				t.Errorf("ParsePosition expected error: %v, got: %v", tc.err, err)
			}
			if out.X != tc.out.X {
				t.Errorf("ParsePosition expected X: %v, got: %v", tc.out.X, out.X)
			}
			if out.Y != tc.out.Y {
				t.Errorf("ParsePosition expected Y: %v, got: %v", tc.out.Y, out.Y)
			}
			if out.Direction != tc.out.Direction {
				t.Errorf("ParsePosition expected Direction: %v, got: %v", tc.out.Direction, out.Direction)
			}
		})
	}
}

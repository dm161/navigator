package navigator

import (
	"fmt"
	"testing"

	"diego.pizza/equalexperts/mars/internal/command"
	"diego.pizza/equalexperts/mars/internal/position"
)

func TestNavigate(t *testing.T) {
	t.Parallel()
	for name, tc := range []struct {
		start  position.Position
		cmd    command.Command
		finish position.Position
	}{
		{
			start:  position.Position{},
			cmd:    command.Command{},
			finish: position.Position{},
		},
		{
			start:  position.Position{
				X: 1,
				Y: 1,
				Direction: position.DirectionNorth,
			},
			cmd:    command.Command{},
			finish: position.Position{
				X: 1,
				Y: 1,
				Direction: position.DirectionNorth,
			},
		},
		{
			start:  position.Position{
				X: 1,
				Y: 1,
				Direction: position.DirectionNorth,
			},
			cmd:    command.Command{
				Sequence: []rune{'F', 'L'},
			},
			finish: position.Position{
				X: 1,
				Y: 2,
				Direction: position.DirectionWest,
			},
		},
		{
			start:  position.Position{
				X: 1,
				Y: 1,
				Direction: position.DirectionNorth,
			},
			cmd:    command.Command{
				Sequence: []rune{'F', 'L', 'B', 'R', 'F'},
			},
			finish: position.Position{
				X: 2,
				Y: 3,
				Direction: position.DirectionNorth,
			},
		},
	} {
		t.Run(fmt.Sprintf("tc-%d", name), func(t *testing.T) {
			out := Navigate(tc.start, tc.cmd)
			if out.X != tc.finish.X {
				t.Errorf("Navigate expected X: %v, got: %v", tc.finish.X, out.X)
			}
			if out.Y != tc.finish.Y {
				t.Errorf("Navigate expected Y: %v, got: %v", tc.finish.Y, out.Y)
			}
			if out.Direction != tc.finish.Direction {
				t.Errorf("Navigate expected Direction: %v, got: %v", tc.finish.Direction, out.Direction)
			}
		})
	}
}

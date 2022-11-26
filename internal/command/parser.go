package command

import "errors"

var (
	ErrCommandCannotBeEmpty = errors.New("command string cannot be empty")
	ErrCommandInvalidValue  = errors.New("invalid command value provided, valid values are: F,B,L,R")
)

func ParseCommand(cmd string) (Command, error) {
	var command Command
	if len(cmd) < 1 {
		return command, ErrCommandCannotBeEmpty
	}
	command.Sequence = make([]rune, 0, len(cmd))
	for _, c := range cmd {
		switch c {
		case 'F', 'B', 'L', 'R':
			command.Sequence = append(command.Sequence, c)
		default:
			return command, ErrCommandInvalidValue
		}
	}
	return command, nil
}

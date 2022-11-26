package main

import (
	"flag"
	"io"
	"log"
	"os"

	"diego.pizza/equalexperts/mars/internal/command"
	"diego.pizza/equalexperts/mars/internal/navigator"
	"diego.pizza/equalexperts/mars/internal/position"
)

func main() {
	var start, command string
	flag.StringVar(&start, "start", "", "Starting position as defined as comma separated string: x,y,direction eg. (1, 2, NORTH)")
	flag.StringVar(&command, "command", "", "Command string as defined by a series of movements: FBLR eg. FFFLRBLFF")
	flag.Parse()
	if command == "" {
		flag.Usage()
		os.Exit(1)
	}
	if start == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(start, command, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(start, cmd string, out io.Writer) error {
	startingPosition, err := position.ParsePosition(start)
	if err != nil {
		return err
	}
	command, err := command.ParseCommand(cmd)
	if err != nil {
		return err
	}
	finalPosition := navigator.Navigate(startingPosition, command)
	out.Write([]byte(finalPosition.ToString()))
	return nil
}

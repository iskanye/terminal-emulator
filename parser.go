package main

import (
	"fmt"
	"os"
	"strings"
	"terminal-emulator/programs"
	"time"
)

// Парсер
func Parser(input string) error {
	parsed := strings.Split(input, " ")
	if parsed[0] == "exit" {
		exit()
	} else {
		program, exists := programs.Programs[parsed[0]]
		if !exists {
			return fmt.Errorf("unknown command")
		} else {
			execute(program, parsed[1:])
		}
	}
	return nil
}

func exit() {
	fmt.Println("Farewell!")
	time.Sleep(500 * time.Millisecond)
	os.Exit(0)
}

func execute(program programs.Program, params []string) {
	out := make(chan interface{})
	in := make(chan interface{})
	execFunc := programFunc(program)

	go transferInput(in, params)
	go execFunc(in, out)

	for i := range out {
		fmt.Print(i)
		fmt.Print(" ")
	}

	fmt.Println()
}

func transferInput(out chan interface{}, input []string) {
	for _, i := range input {
		out <- i
	}
	close(out)
}

func programFunc(program programs.Program) programs.Program {
	return func(in, out chan interface{}) {
		program(in, out)
		close(out)
	}
}

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
			return fmt.Errorf("unknown command: %s", parsed[0])
		} else {
			return execute(program, parsed[1:])
		}
	}
	return nil
}

func exit() {
	fmt.Println("Farewell!")
	time.Sleep(500 * time.Millisecond)
	os.Exit(0)
}

func execute(program programs.Program, params []string) error {
	stdin := make(chan string)
	stdout := make(chan interface{})
	stderr := make(chan error)
	execFunc := programFunc(program)

	// Поток передачи аргументов
	go func() {
		for _, i := range params {
			stdin <- i
		}
		close(stdin)
	}()

	go execFunc(stdin, stdout, stderr)

	// Вывод программы
	go func() {
		for i := range stdout {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}()

	// Обработка ошибок
	if err := <-stderr; err != nil {
		return err
	}

	return nil
}

// Поток программы
func programFunc(program programs.Program) programs.Program {
	return func(in chan string, out chan interface{}, err chan error) {
		program(in, out, err)
		close(out)
		close(err)
	}
}

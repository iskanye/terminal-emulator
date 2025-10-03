package main

import (
	"fmt"
	"os"
	"strings"

	"terminal-emulator/programs"
	"terminal-emulator/vfs"
)

// Парсер
func Parser(input string) error {
	parsed := strings.Split(input, " ")

	if parsed[0] == "exit" {
		exit()
		return nil
	} else {
		program, exists := programs.Programs[parsed[0]]
		if !exists {
			return fmt.Errorf("unknown command: %s", parsed[0])
		} else {
			return execute(program, parsed[1:])
		}
	}
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

	// Поток самой программы
	go execFunc(stdin, stdout, stderr)

	// Обработка ошибок
	select {
	case err := <-stderr:
		return err
	default:
		// Вывод программы
		for i := range stdout {
			Println(i)
		}
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

func exit() {
	vfs.FileExplorer.Save(vfsPath)
	os.Exit(0)
}

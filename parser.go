package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

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
	wg := sync.WaitGroup{}
	stdin := make(chan string)
	stdout := make(chan any)
	stderr := make(chan error)
	execFunc := programFunc(program)
	wg.Add(3)

	// Поток передачи аргументов
	go func() {
		defer wg.Done()
		for _, i := range params {
			stdin <- i
		}
		close(stdin)
	}()

	go func() {
		defer wg.Done()
		execFunc(stdin, stdout, stderr)
	}()

	// Вывод программы
	go func() {
		defer wg.Done()

		for i := range stdout {
			Println(i)
		}
	}()

	// Обработка ошибок
	if err := <-stderr; err != nil {
		return err
	}

	// Ожидание завершения всех потоков
	wg.Wait()

	return nil
}

// Поток программы
func programFunc(program programs.Program) programs.Program {
	return func(in chan string, out chan any, err chan error) {
		program(in, out, err)
		close(out)
		close(err)
	}
}

func exit() {
	vfs.FileExplorer.Save(vfsPath)
	os.Exit(0)
}

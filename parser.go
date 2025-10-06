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

func execute(program func(), params []string) error {
	wg := sync.WaitGroup{}
	programs.InitChannels()
	wg.Add(3)

	// Поток передачи аргументов
	go func() {
		defer wg.Done()
		programs.WriteToStdin(params)
	}()

	go func() {
		defer wg.Done()
		program()
	}()

	// Вывод программы
	go func() {
		defer wg.Done()

		for i := range programs.Stdout() {
			Println(i)
		}
	}()

	// Обработка ошибок
	if err := <-programs.Stderr(); err != nil {
		return err
	}

	// Ожидание завершения всех потоков
	wg.Wait()

	return nil
}

func exit() {
	vfs.FileExplorer.Save(vfsPath)
	os.Exit(0)
}

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

func execute(program func(), params []string) error {
	return programs.Execute(program, params, func(i any) {
		Println(i)
	})
}

func exit() {
	vfs.FileExplorer.Save(vfsPath)
	os.Exit(0)
}

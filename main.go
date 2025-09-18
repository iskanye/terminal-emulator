package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"terminal-emulator/vfs"
)

var (
	username    = "iskanye"
	vfsPath     = "root.xml"
	startScript = "start"
)

// Поле ввода
func PrintInputField() {
	fmt.Print(username + ":" + vfs.FileExplorer.GetPosition() + "> ")
}

func main() {
	if len(os.Args) > 1 {
		vfsPath = os.Args[1]
		startScript = os.Args[2]
	}

	err := setupVFS()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Welcome to terminal emulator! (~by iskanye~)\n" +
		"VFS: " + vfsPath + "\n" +
		"Script: " + startScript)

	ExecuteScript(startScript)
	terminal()
}

// Основной цикл эмулятора
func terminal() {
	reader := bufio.NewReader(os.Stdin)

	for {
		PrintInputField()

		input, _ := reader.ReadString('\n')
		err := Parser(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func setupVFS() error {
	fs, err := vfs.LoadFromXML(vfsPath)
	if err != nil {
		return err
	}
	vfs.SetupExplorer(fs)
	return nil
}

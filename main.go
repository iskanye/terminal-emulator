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

var fileSystem *vfs.Node

func main() {
	if len(os.Args) > 1 {
		vfsPath = os.Args[1]
		startScript = os.Args[2]
	}

	setupVFS()
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

func PrintInputField() {
	fmt.Print(username + "> ")
}

func setupVFS() {
	fs := vfs.NewRoot()

	// Создаем директории и файлы
	fs.Create("/home", true)
	fs.Create("/home/user", true)
	fs.Create("/home/user/document.txt", false)
	fs.Create("/home/user/image.jpg", false)
	fs.Create("/etc", true)
	fs.Create("/etc/config.conf", false)

	err := fs.SaveToXML(vfsPath)
	if err != nil {
		fmt.Println(err)
	}

	fileSystem, err = vfs.LoadFromXML(vfsPath)
	if err != nil {
		fmt.Println(err)
	}
}

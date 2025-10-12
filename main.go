package main

import (
	"fmt"
	"os"
	"strings"

	"terminal-emulator/vfs"

	"gioui.org/app"
)

type reader interface {
	Read() (string, error)
}

var (
	username    = "iskanye"
	vfsPath     = "root.xml"
	startScript = "start"
)

var terminal *Terminal

func Print(a any) {
	terminal.Print(a)
}

func Println(a any) {
	terminal.Println(a)
}

// Поле ввода
func PrintInputField() {
	Print(username + ":" + vfs.FileExplorer.GetPosition() + "> ")
}

func setupVFS() {
	fs := vfs.LoadFromXML(vfsPath)
	vfs.SetupExplorer(fs)
}

func main() {
	if len(os.Args) > 1 {
		vfsPath = os.Args[1]
		startScript = os.Args[2]
	}

	terminal = NewTerminal("VFS: " + vfsPath)
	go terminal.Main()

	setupVFS()
	Println("Welcome to terminal emulator! (~by iskanye~)\n" +
		"VFS: " + vfsPath + "\n" +
		"Script: " + startScript)

	var reader reader = NewScript(startScript)

	go func() {
		for {
			text, err := reader.Read()
			if err != nil {
				reader = terminal
				continue
			}

			text = strings.TrimSpace(text)
			Println(text)

			err = Parser(text)
			if err != nil {
				switch r := reader.(type) {
				case *Script:
					Println(fmt.Sprintf("line %d: ", r.CurrentLine) + " " + fmt.Sprint(err))
					reader = terminal
				case *Terminal:
					Println(err)
				}
			}

			PrintInputField()
		}
	}()

	app.Main()
}

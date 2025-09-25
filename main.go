package main

import (
	"os"
	"terminal-emulator/vfs"

	"gioui.org/app"
)

var (
	username    = "iskanye"
	vfsPath     = "root.xml"
	startScript = "start"
)

var terminal *Terminal

func Print(a interface{}) {
	terminal.Print(a)
}

func Println(a interface{}) {
	terminal.Println(a)
}

// Поле ввода
func PrintInputField() {
	Print(username + ":" + vfs.FileExplorer.GetPosition() + "> ")
}

func main() {
	if len(os.Args) > 1 {
		vfsPath = os.Args[1]
		startScript = os.Args[2]
	}

	terminal = NewTerminal("VFS: " + vfsPath)
	go terminal.Main()

	err := setupVFS()
	if err != nil {
		Print(err)
		return
	}

	Println("Welcome to terminal emulator! (~by iskanye~)\n" +
		"VFS: " + vfsPath + "\n" +
		"Script: " + startScript)

	go ExecuteScript(startScript)

	go func() {
		for {
			text := terminal.Read()
			Println(text)

			err := Parser(text)
			if err != nil {
				Println(err)
			}

			PrintInputField()
		}
	}()

	app.Main()
}

func setupVFS() error {
	fs, err := vfs.LoadFromXML(vfsPath)
	if err != nil {
		return err
	}
	vfs.SetupExplorer(fs)
	return nil
}

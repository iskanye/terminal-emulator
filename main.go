package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var username string = "iskanye"
var vfs string = ""
var startScript string = "start"

func main() {
	vfs = os.Args[1]
	startScript = os.Args[2]

	fmt.Println("Welcome to terminal emulator! (~by iskanye~)\n" +
		"VFS: " + vfs + "\n" +
		"Script: " + startScript)

	ExecuteScript(startScript)
	terminal()
}

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

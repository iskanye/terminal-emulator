package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Cd(in chan string, out chan interface{}, err chan error) {
	args := ExtractArgs(in)

	if len(args) == 0 {
		err <- fmt.Errorf("no args")
		return
	} else if len(args) > 1 {
		err <- fmt.Errorf("too many arguments")
		return
	}

	err <- vfs.FileExplorer.Travel(args[0])
}

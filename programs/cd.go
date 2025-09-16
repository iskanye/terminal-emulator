package programs

import (
	"fmt"
	"terminal-emulator/vfs"
)

func Cd(in chan string, out chan interface{}, err chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	if len(args) == 0 {
		err <- fmt.Errorf("no args")
		return
	}

	err <- vfs.FileExplorer.Travel(args[0])
}

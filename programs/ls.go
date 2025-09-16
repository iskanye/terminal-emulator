package programs

import (
	"fmt"
	"terminal-emulator/vfs"
)

func Ls(in chan string, out chan interface{}, err chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	for _, i := range args {
		err <- fmt.Errorf("unknown argument: %s", i)
	}

	result := vfs.FileExplorer.List()
	for _, i := range result {
		out <- i
	}

	err <- nil
}

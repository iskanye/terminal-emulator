package programs

import (
	"terminal-emulator/vfs"
)

func Cd(in chan string, out chan interface{}, err chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	err <- vfs.FileExplorer.Travel(args[0])
}

package programs

import (
	"fmt"
	"terminal-emulator/vfs"
)

func Mkdir(in chan string, out chan interface{}, stderr chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	}
	if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	err := vfs.FileExplorer.AddNode(args[0], false)
	if err != nil {
		stderr <- err
		return
	}

	stderr <- nil
}

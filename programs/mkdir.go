package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Mkdir(in chan string, out chan any, stderr chan error) {
	args := ExtractArgs(in)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	}
	if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	err := vfs.FileExplorer.AddNode(args[0], true)
	if err != nil {
		stderr <- err
		return
	}

	stderr <- nil
}

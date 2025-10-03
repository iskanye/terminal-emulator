package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Cd(in chan string, out chan interface{}, stderr chan error) {
	args := ExtractArgs(in)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	} else if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	if err := vfs.FileExplorer.Travel(args[0]); err != nil {
		stderr <- err
	}
}

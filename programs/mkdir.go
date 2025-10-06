package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Mkdir() {
	args := ExtractArgs(stdin)

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

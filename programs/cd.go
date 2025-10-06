package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Cd() {
	args := ExtractArgs(stdin)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	} else if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	stderr <- vfs.FileExplorer.Travel(args[0])
}

package programs

import (
	"fmt"

	"terminal-emulator/vfs"
)

func Rmdir(in chan string, out chan interface{}, stderr chan error) {
	args := ExtractArgs(in)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	}
	if len(args) > 1 {
		stderr <- fmt.Errorf("too many arguments")
		return
	}

	node, err := vfs.FileExplorer.GetNode(args[0])
	if err != nil {
		stderr <- err
		return
	}

	if node == vfs.FileExplorer.GetCurrent() {
		stderr <- fmt.Errorf("cannot delete current directory")
		return
	}
	if !node.IsDirectory {
		stderr <- fmt.Errorf("%s isn`t directory", args[0])
		return
	}
	if len(node.Children) != 0 {
		stderr <- fmt.Errorf("directory isn`t empty")
		return
	}

	err = node.Delete()
	if err != nil {
		stderr <- err
		return
	}

	stderr <- nil
}

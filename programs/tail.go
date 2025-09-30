package programs

import (
	"fmt"
	"strconv"
	"strings"

	"terminal-emulator/vfs"
)

func Tail(in chan string, out chan interface{}, stderr chan error) {
	n := 10
	args := ExtractArgs(in)

	if len(args) == 0 {
		stderr <- fmt.Errorf("no args")
		return
	} else if len(args) > 1 && len(args) < 4 {
		argv, err := ExtractArgv(args[:len(args)-1])
		if err != nil {
			stderr <- err
			return
		}

		if val, ok := argv["n"]; ok {
			n, err = strconv.Atoi(val)
			if err != nil {
				stderr <- err
				return
			}
		} else if val, ok := argv["lines"]; ok {
			n, err = strconv.Atoi(val)
			if err != nil {
				stderr <- err
				return
			}
		} else {
			stderr <- fmt.Errorf("wrong arg: %s", args[0])
			return
		}
	} else if len(args) != 1 {
		stderr <- fmt.Errorf("too many args")
		return
	}

	file, err := vfs.FileExplorer.GetFile(args[len(args)-1])
	if err != nil {
		stderr <- err
		return
	}

	content, err := file.Read()
	if err != nil {
		stderr <- err
		return
	}

	if content != "" {
		lines := strings.Split(content, "\n")

		if len(lines) <= n {
			out <- strings.TrimSpace(content)
		} else {
			for i := n; i > 0; i-- {
				out <- lines[len(lines)-i]
			}
		}
	}

	stderr <- nil
}

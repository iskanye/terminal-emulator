package programs

import (
	"bufio"
	"os"
	"strings"
	"terminal-emulator/vfs"
)

func Cat(in chan string, out chan interface{}, stderr chan error) {
	args := make([]string, 0)
	for i := range in {
		args = append(args, i)
	}

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		result, _ := reader.ReadString('\n')
		out <- strings.TrimSpace(result)
	} else {
		for _, i := range args {
			file, err := vfs.FileExplorer.GetFile(i)
			if err != nil {
				stderr <- err
				return
			}

			content, err := file.Read()
			if err != nil {
				stderr <- err
				return
			}

			out <- strings.TrimSpace(content)
		}
	}

	stderr <- nil
}

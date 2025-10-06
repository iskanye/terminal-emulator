// Пакет со всеми доступными встроенными командами
package programs

import (
	"fmt"
	"strings"
)

// Входной канал в который поступают аргументы программы;
var stdin chan string

// Записывает данные во входной канал
func WriteToStdin(data []string) {
	for _, i := range data {
		stdin <- i
	}
	close(stdin)
}

// Выходной канал в который поступает результат работы программы;
var stdout chan any

func Stdout() chan any {
	return stdout
}

// Канал исключений, при успешном выполнении в него поступает nil
var stderr chan error

func Stderr() chan error {
	return stderr
}

// Инициализировать основные каналы программ
func InitChannels() {
	stdin = make(chan string)
	stdout = make(chan any)
	stderr = make(chan error)
}

// Встроенные команды
var Programs = map[string]func(){
	"ls":    Program(Ls),
	"cd":    Program(Cd),
	"du":    Program(Du),
	"tail":  Program(Tail),
	"cat":   Program(Cat),
	"touch": Program(Touch),
	"rmdir": Program(Rmdir),
	"mkdir": Program(Mkdir),
	"help":  Program(Help),
	"pico":  Program(Pico),
}

func Program(programFunc func()) func() {
	return func() {
		programFunc()
		close(stdout)
		close(stderr)
	}
}

// Получет аргументы из канала
func ExtractArgs(stdin chan string) []string {
	args := make([]string, 0)
	for i := range stdin {
		args = append(args, i)
	}

	return args
}

// Функция извлечения именованных параметров
func ExtractArgv(args []string) (map[string]string, error) {
	if len(args)%2 != 0 {
		return nil, fmt.Errorf("wrong number of args")
	}

	result := make(map[string]string)
	for i := 0; i < len(args); i++ {
		key := strings.Trim(args[i], "-")
		i++
		result[key] = args[i]
	}

	return result, nil
}

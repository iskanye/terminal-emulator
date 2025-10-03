// Пакет со всеми доступными встроенными командами
package programs

import (
	"fmt"
	"strings"
)

// Исполняемая функции программы:
// in - входной канал в который поступают аргументы программы;
// out - выходной канал в который поступает результат работы программы;
// err - канал исключений, при успешном выполнении в него поступает nil
type Program func(in chan string, out chan interface{}, err chan error)

// Встроенные команды
var Programs = map[string]Program{
	"ls":    Ls,
	"cd":    Cd,
	"du":    Du,
	"tail":  Tail,
	"cat":   Cat,
	"touch": Touch,
	"rmdir": Rmdir,
	"mkdir": Mkdir,
	"help":  Help,
	"pico":  Pico,
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

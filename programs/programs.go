// Пакет со всеми доступными встроенными командами
package programs

import (
	"fmt"
	"strings"
	"sync"
)

// Входной канал в который поступают аргументы программы;
var stdin chan string

// Выходной канал в который поступает результат работы программы;
var stdout chan any

// Канал исключений, при успешном выполнении в него поступает nil
var stderr chan error

// Встроенные команды
var Programs = map[string]func(){
	"ls":    program(Ls),
	"cd":    program(Cd),
	"du":    program(Du),
	"tail":  program(Tail),
	"cat":   program(Cat),
	"touch": program(Touch),
	"rmdir": program(Rmdir),
	"mkdir": program(Mkdir),
	"help":  program(Help),
	"pico":  program(Pico),
}

// Выполнить программу
func Execute(program func(), input []string, handleOutput func(any)) error {
	// Инициализировать основные каналы программ
	stdin = make(chan string)
	stdout = make(chan any)
	stderr = make(chan error)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	// Поток передачи аргументов
	go func() {
		defer wg.Done()
		for _, i := range input {
			stdin <- i
		}
		close(stdin)
	}()

	// Поток программы
	go func() {
		defer wg.Done()
		program()
	}()

	// Вывод программы
	go func() {
		defer wg.Done()
		for i := range stdout {
			handleOutput(i)
		}
	}()

	// Обработка ошибок
	if err := <-stderr; err != nil {
		return err
	}

	// Ожидание завершения всех потоков
	wg.Wait()

	return nil
}

func program(programFunc func()) func() {
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

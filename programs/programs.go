// Пакет со всеми доступными встроенными командами
package programs

type Program func(in, out chan interface{})

// Встроенные команды
var Programs = map[string]Program{
	"ls": Ls,
	"cd": Cd,
}

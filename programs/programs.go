package programs

type Program func(in, out chan interface{})

var Programs = map[string]Program{
	"ls": Ls,
	"cd": Cd,
}

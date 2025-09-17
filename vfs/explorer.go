package vfs

import "fmt"

type Explorer struct {
	current *Node
}

// Экземпляр просмотрщика файловой системы
var FileExplorer Explorer

// Запуск эксплорера
func SetupExplorer(root *Node) {
	FileExplorer = Explorer{
		current: root,
	}
}

// Перемещение по файловой системе
func (exp *Explorer) Travel(path string) error {
	if path == "." {
		return nil
	}
	if path == ".." && exp.current.Parent == nil {
		return fmt.Errorf("can`t move higher from root")
	}
	if path == ".." {
		exp.current = exp.current.Parent
		return nil
	}

	// Возвращаемся к корневой ноде
	if path[0] == '/' {
		exp.returnToRoot()
	}

	node, err := exp.current.GetNode(path)
	if err != nil {
		return err
	}

	if !node.IsDirectory {
		return fmt.Errorf("can`t move to file")
	}

	exp.current = node
	return nil
}

// Возвращает файл из текущей директории
func (exp *Explorer) GetFile(name string) (*Node, error) {
	for _, i := range exp.current.Children {
		if i.Name == name {
			return i, nil
		}
	}

	return nil, fmt.Errorf("can`t find file: %s", name)
}

// Получает текстовый список нод в текущей директории
func (exp *Explorer) List() []string {
	list := make([]string, 0)

	for _, i := range exp.current.Children {
		list = append(list, i.Name)
	}

	return list
}

// Получает текстовый список нод в данной директории
func (exp *Explorer) ListDir(path string) ([]string, error) {
	dir, err := exp.current.GetNode(path)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)

	for _, i := range dir.Children {
		list = append(list, i.Name)
	}

	return list, nil
}

// Получает позицию текущей ноды относительно корневой ноды
func (exp *Explorer) GetPosition() string {
	if exp.current.Parent == nil {
		return "/"
	}

	position := ""
	current := exp.current
	for current.Parent != nil {
		position = "/" + current.Name + position
		current = current.Parent
	}

	return position
}

// Текущая нода
func (exp *Explorer) GetCurrent() *Node {
	return exp.current
}

// Возвращается в корневую ноду
func (exp *Explorer) returnToRoot() {
	for exp.current.Parent != nil {
		exp.current = exp.current.Parent
	}
}

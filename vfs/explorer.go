package vfs

import (
	"fmt"
	"strings"
	"time"
)

type Explorer struct {
	current *Node
	root    *Node
}

// Экземпляр просмотрщика файловой системы
var FileExplorer Explorer

// Запуск эксплорера
func SetupExplorer(root *Node) {
	FileExplorer = Explorer{
		current: root,
		root:    root,
	}
}

// Перемещение по файловой системе
func (exp *Explorer) Travel(path string) error {
	if path == "." || path == "" {
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
		if i.Name == name && i.IsDirectory {
			return nil, fmt.Errorf("not a file: %s", name)
		} else if i.Name == name && !i.IsDirectory {
			return i, nil
		}
	}

	return nil, fmt.Errorf("can`t find file: %s", name)
}

// Возвращает ноду в текущей файловой системе
func (exp *Explorer) GetNode(path string) (*Node, error) {
	current := exp.current
	if path[0] == '/' {
		current = exp.root
	}

	return current.GetNode(strings.Trim(path, "/"))
}

// Создать ноду в текущей директории
func (exp *Explorer) AddNode(name string, isDir bool) error {
	// Проверяем, существует ли уже элемент с таким именем
	for _, child := range exp.current.Children {
		if child.Name == name {
			return fmt.Errorf("file or directory already exists: %s", name)
		}
	}

	node := &Node{
		Name:        name,
		Parent:      exp.current,
		IsDirectory: isDir,
		Modified:    time.Now(),
	}
	if isDir {
		node.Children = make([]*Node, 0)
	}

	exp.current.Children = append(exp.current.Children, node)

	return nil
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

	if !dir.IsDirectory {
		return nil, fmt.Errorf("%s is a file", path)
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

// Сохранить файловую систему в файл
func (exp *Explorer) Save(path string) {
	exp.root.SaveToXML(path)
}

// Возвращается в корневую ноду
func (exp *Explorer) returnToRoot() {
	exp.current = exp.root
}

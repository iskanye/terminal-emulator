package vfs

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

// Создать новую корневую ветку
func NewRoot() *Node {
	root := Node{
		Name:        "",
		IsDirectory: true,
		Children:    make([]*Node, 0),
		Parent:      nil,
	}
	return &root
}

// Получить указатель на ноду по её пути в данной ветке
func (root *Node) GetNode(path string) (*Node, error) {
	if path == "" {
		return root, nil
	}

	// Разделяем путь
	parts := strings.Split(path, "/")

	current := root
	for _, i := range parts {
		if i == "" {
			continue
		}

		found := false
		for _, child := range current.Children {
			if child.Name == i {
				current = child
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("path not found: %s", path)
		}
	}

	return current, nil
}

// Удалить данную ноду
func (root *Node) Delete() error {
	parent := root.Parent
	if parent == nil {
		return fmt.Errorf("cannot delete root directory")
	}

	// Удаляем узел из родительского списка детей
	for i, child := range parent.Children {
		if child == root {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return nil
		}
	}

	return nil
}

// Записать строку в ноду
func (root *Node) Write(data string) error {
	if root.IsDirectory {
		return fmt.Errorf("cannot write to directory")
	}

	root.Content = data
	root.Modified = time.Now()
	return nil
}

// Записать байты в ноду
func (root *Node) WriteBytes(data []byte) error {
	return root.Write(base64.StdEncoding.EncodeToString(data))
}

// Прочитать данные с ноды
func (root *Node) Read() (string, error) {
	if root.IsDirectory {
		return "", fmt.Errorf("cannot read from directory")
	}

	return root.Content, nil
}

// Прочитать байты с ноды
func (root *Node) ReadBytes() ([]byte, error) {
	if root.IsDirectory {
		return nil, fmt.Errorf("cannot read from directory")
	}

	data, err := base64.RawStdEncoding.DecodeString(root.Content)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Размер ноды в байтах
func (root *Node) GetSize() int {
	if !root.IsDirectory {
		return len(root.Content)
	}

	size := 0
	for _, i := range root.Children {
		size += i.GetSize()
	}
	return size
}

// Восстановить указатели на родителей в данной ветке
func (root *Node) restoreParent(parent *Node) {
	root.Parent = parent

	for _, i := range root.Children {
		i.restoreParent(root)
	}
}

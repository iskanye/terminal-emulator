package vfs

import (
	"fmt"
	"path/filepath"
	"strings"
)

func NewFileSystem() *Node {
	root := Node{
		Name:        "/",
		IsDirectory: true,
		Children:    make([]*Node, 0),
	}
	return &root
}

func (root *Node) GetNode(path string) (*Node, error) {
	if path == "/" {
		return root, nil
	}

	// Нормализуем путь
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")

	current := root
	for _, part := range parts {
		if part == "" {
			continue
		}

		found := false
		for _, child := range current.Children {
			if child.Name == part {
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

func (root *Node) Create(path string, isDir bool) error {
	dir, name := filepath.Split(path)
	parent, err := root.GetNode(dir)
	if err != nil {
		return fmt.Errorf("parent directory not found: %v", err)
	}

	if parent.IsDirectory {
		return fmt.Errorf("%s isn`t a directory", name)
	}

	// Проверяем, существует ли уже элемент с таким именем
	for _, child := range parent.Children {
		if child.Name == name {
			return fmt.Errorf("file or directory already exists: %s", name)
		}
	}

	node := &Node{
		Name:        name,
		Parent:      parent,
		IsDirectory: isDir,
	}
	if isDir {
		node.Children = make([]*Node, 0)
	}

	parent.Children = append(parent.Children, node)

	return nil
}

func (root *Node) Delete(path string) error {
	if path == "/" {
		return fmt.Errorf("cannot delete root directory")
	}

	node, err := root.GetNode(path)
	if err != nil {
		return err
	}

	parent := node.Parent

	// Удаляем узел из родительского списка детей
	for i, child := range parent.Children {
		if child == node {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("node not found in parent's children")
}

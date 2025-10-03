package vfs

import (
	"encoding/xml"
	"os"
)

// Загрузить корневую ноду с xml файла
func LoadFromXML(file string) *Node {
	xmlData, err := os.ReadFile(file)
	if err != nil {
		_, _ = os.Create(file)
		return NewRoot()
	}

	root := &Node{}
	if err := xml.Unmarshal([]byte(xmlData), root); err != nil {
		root := NewRoot()
		root.SaveToXML(file)
		return NewRoot()
	}

	// Создаём в каждой ноде указатель на родителя
	root.restoreParent(nil)

	return root
}

// Сохранить ветку в xml файл
func (root *Node) SaveToXML(filename string) error {
	data, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(string(data)); err != nil {
		return err
	}

	return nil
}

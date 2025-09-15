package vfs

import (
	"encoding/xml"
	"os"
)

func LoadFromXML(file string) (*Node, error) {
	xmlData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	root := &Node{}
	if err := xml.Unmarshal([]byte(xmlData), root); err != nil {
		return nil, err
	}

	return root, nil
}

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

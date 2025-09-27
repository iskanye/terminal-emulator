package main

import (
	"bufio"
	"os"
)

// Тип скрипта
type Script struct {
	reader      *bufio.Reader
	CurrentLine int
}

// Загрузить скрипт
func NewScript(script string) *Script {
	file, err := os.Open(script)

	if err != nil {
		Println(err)
		return nil
	}

	return &Script{
		reader:      bufio.NewReader(file),
		CurrentLine: 0,
	}
}

// Читать строку скрипта
func (s *Script) Read() (string, error) {
	s.CurrentLine++
	return s.reader.ReadString('\n')
}

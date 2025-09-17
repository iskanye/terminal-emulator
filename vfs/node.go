package vfs

import (
	"encoding/xml"
	"time"
)

type Node struct {
	XMLName     xml.Name  `xml:"node"`
	Name        string    `xml:"name,attr"`
	IsDirectory bool      `xml:"dir,attr"`
	Modified    time.Time `xml:"modified,attr"`
	Content     string    `xml:"content,omitempty"`
	Children    []*Node   `xml:"node"`
	Parent      *Node     `xml:"-"`
}

package vfs

import "encoding/xml"

type Node struct {
	XMLName     xml.Name `xml:"node"`
	Name        string   `xml:"name,attr"`
	IsDirectory bool     `xml:"isDirectory,attr"`
	Content     string   `xml:"content,omitempty"`
	Children    []*Node  `xml:"node"`
	Parent      *Node    `xml:"-"`
}

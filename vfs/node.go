package vfs

type Node struct {
	Name        string  `xml:"node"`
	IsDirectory bool    `xml:"isDirectory,attr"`
	Content     string  `xml:"content,omitempty"`
	Children    []*Node `xml:"children>node,omitempty"`
	Parent      *Node   `xml:"-"`
}

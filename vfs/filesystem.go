package vfs

import "encoding/xml"

type FileSystem struct {
	Name xml.Name `xml:"filesystem"`
	Root *Node    `xml:"root"`
}

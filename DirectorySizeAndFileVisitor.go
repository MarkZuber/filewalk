package main

import (
	"fmt"
	"path/filepath"
)

type DirectorySizeAndFileVisitor struct {
}

func (dnv *DirectorySizeAndFileVisitor) Visit(dn *DirectoryNode, indentLevel int) {
	dsv := &DirectorySizeVisitor{}
	dsv.Visit(dn, indentLevel)

	indentString := getIndentString(indentLevel + 1)
	colorStart := colorBlue
	colorEnd := colorFinish

	for path, fileInfo := range dn.fileInfos {
		fileName := filepath.Base(path)
		fileSize := formatBytes(fileInfo.Size())
		fmt.Printf("%s%s%s: %s%s\n", colorStart, indentString, fileName, fileSize, colorEnd)
	}
}

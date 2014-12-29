package main

import (
	// "fmt"
	"os"
	"path/filepath"
)

type DirectoryNode struct {
	fullPath    string
	localSize   int64
	totalSize   int64
	fileInfos   map[string]os.FileInfo
	subDirNodes []*DirectoryNode
}

type ProgressCallbackFunc func(fullPath string)

type DirectoryNodeVisitor interface {
	Visit(dn *DirectoryNode, indentLevel int)
}

// func (dn *DirectoryNode) accept(visitor *DirectorySizeVisitor, indentLevel int) {
func (dn *DirectoryNode) Accept(visitor DirectoryNodeVisitor, indentLevel int) {
	visitor.Visit(dn, indentLevel)
	for _, subdir := range dn.subDirNodes {
		subdir.Accept(visitor, indentLevel+1)
	}
}

func (dn *DirectoryNode) scan(progressCallback ProgressCallbackFunc) {

	// fmt.Printf("scan entered (%s)\n", dn.fullPath)
	// progressCallback(dn.fullPath)

	// initialize the map
	dn.fileInfos = make(map[string]os.FileInfo)

	filepath.Walk(dn.fullPath, func(path string, info os.FileInfo, err error) error {
		// fmt.Println(path)

		if info.IsDir() {
			if path == dn.fullPath {
				// fmt.Printf("we're in the same dir %s\n", path)
				return nil
			} else {
				// fmt.Printf("adding a subdir node for %s\n", path)
				dn.subDirNodes = append(dn.subDirNodes, &DirectoryNode{fullPath: path})
				return filepath.SkipDir
			}
		} else {
			// fmt.Printf("we're a file: %s (%s) (%s)\n", path, filepath.Dir(path), dn.fullPath)
			if filepath.Dir(path) == dn.fullPath {
				// fmt.Println("adding file info")
				dn.fileInfos[path] = info
				dn.localSize += info.Size()
			}
			return nil
		}
	})

	dn.totalSize += dn.localSize

	for _, subdir := range dn.subDirNodes {
		subdir.scan(progressCallback)
		dn.totalSize += subdir.totalSize
	}
}

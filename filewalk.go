package main

import (
	"fmt"
)

type FileWalker struct {
	rootDirectory string
	rootDirNode   DirectoryNode
	// lcv LineCountVisitor
}

func progressCallback(fullPath string) {
	fmt.Println(fullPath)
}

func (fw *FileWalker) Build() {
	fw.rootDirectory = "/Users/zube/Downloads"

	fw.rootDirNode = DirectoryNode{fullPath: fw.rootDirectory}
	fw.rootDirNode.scan(progressCallback)
}

func (fw *FileWalker) PrintDirectorySizes() {
	dsv := &DirectorySizeVisitor{}
	fw.rootDirNode.Accept(dsv, 0)
}

func (fw *FileWalker) PrintDirectoryAndFileSizes() {
	dsfv := &DirectorySizeAndFileVisitor{}
	fw.rootDirNode.Accept(dsfv, 0)
}

func main() {
	fmt.Println("FileWalk by Mark Zuber")

	var walker FileWalker
	walker.Build()
	// walker.PrintDirectorySizes()
	walker.PrintDirectoryAndFileSizes()
}

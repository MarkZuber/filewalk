package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

type DirectorySizeVisitor struct {
}

const colorHeader = "\033[30m"
const colorBlue = "\033[34m"
const colorGreen = "\033[36m"
const colorWarning = "\033[37m"
const colorFail = "\033[31m"
const colorFinish = "\033[0m"
const colorBold = "\033[1m"

func (dnv *DirectorySizeVisitor) Visit(dn *DirectoryNode, indentLevel int) {
	dirName := dn.fullPath
	if indentLevel > 0 {
		dirName = filepath.Base(dirName)
	}

	indentString := getIndentString(indentLevel)
	totalBytes := formatBytes(dn.totalSize)
	localBytes := formatBytes(dn.localSize)
	numFiles := fmt.Sprintf("%d", len(dn.fileInfos))

	colorStart := colorGreen
	colorEnd := colorFinish

	if dn.totalSize >= MEGABYTE {
		colorStart = colorWarning
	}
	if dn.totalSize >= GIGABYTE {
		colorStart = colorBold + colorFail
	}

	fmt.Printf(
		"%s%s%s: %s (%s local in %s files)%s\n",
		colorStart,
		indentString,
		dirName,
		totalBytes,
		localBytes,
		numFiles,
		colorEnd)
}

func getIndentString(indentLevel int) string {
	return strings.Repeat(" ", indentLevel*2) // 2 space tabs
}

const KILOBYTE = 1024
const MEGABYTE = 1048576
const GIGABYTE = 1073741824
const TERABYTE = 1099511627776

func formatBytes(numBytes int64) string {
	if numBytes < KILOBYTE {
		return fmt.Sprintf("%d", numBytes)
	} else if numBytes < MEGABYTE {
		return fmt.Sprintf("%.2fKB", float64(numBytes)/float64(KILOBYTE))
	} else if numBytes < GIGABYTE {
		return fmt.Sprintf("%.2fMB", float64(numBytes)/float64(MEGABYTE))
	} else if numBytes < TERABYTE {
		return fmt.Sprintf("%.2fGB", float64(numBytes)/float64(GIGABYTE))
	} else {
		return fmt.Sprintf("%.2fTB", float64(numBytes)/float64(TERABYTE))
	}
}

package main

import (
	"fmt"
	"strings"

	"github.com/wizzymore/tinyfiledialogs"
)

func main() {
	sel, ok := tinyfiledialogs.OpenFileDialog("Source Files", "", []string{"*.go", "*.mod"}, "Go files", false)
	if ok {
		fmt.Println(sel)
	} else {
		fmt.Println("No file selected")
	}

	sel, ok = tinyfiledialogs.OpenFileDialog("Source Files", "", []string{"*.go", "*.mod"}, "Go files", true)
	if ok {
		fmt.Println(strings.Split(sel, "|"))
	} else {
		fmt.Println("No files selected")
	}

	sel, ok = tinyfiledialogs.SelectFolderDialog("Source Files", "")
	if ok {
		fmt.Println(sel)
	} else {
		fmt.Println("No folder selected")
	}

	sel, ok = tinyfiledialogs.SaveFileDialog("Source Files", "", []string{"*.go", "*.mod"}, "Go files")
	if ok {
		fmt.Println(sel)
	} else {
		fmt.Println("No file selected")
	}
}

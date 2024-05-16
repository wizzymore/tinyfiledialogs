package main

import (
	"fmt"
	"os"

	"github.com/wizzymore/tinyfiledialogs"
)

func main() {
	sel, ok := tinyfiledialogs.OpenFileDialog("Source Files", "", []string{"*.go", "*.mod"}, "Go files", false)
	if ok {
		fmt.Println("OpenFileDialog - ", sel)
	} else {
		fmt.Println("OpenFileDialog - No file selected")
	}

	sel, ok = tinyfiledialogs.OpenFileDialog("Source Files", "", []string{"*.go", "*.mod"}, "Go files", true)
	if ok {
		fmt.Println("OpenFileDialog Multiple - ", sel)
	} else {
		fmt.Println("OpenFileDialog Multiple - No files selected")
	}

	sel, ok = tinyfiledialogs.SelectFolderDialog("Source Files", "")
	if ok {
		fmt.Println("SelectFolderDialog - ", sel)
	} else {
		fmt.Println("SelectFolderDialog - No folder selected")
	}

	sel, ok = tinyfiledialogs.SaveFileDialog("Source Files", "", []string{"*.go", "*.mod"}, "Go files")
	if ok {
		fmt.Println("SaveFileDialog - ", sel)
	} else {
		fmt.Println("SaveFileDialog - Action canceled")
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not get home dir")
		return
	}

	sel, ok = tinyfiledialogs.OpenFileDialog("Source Files", homedir+"/", []string{"*.go", "*.mod"}, "Go files", false)
	if ok {
		fmt.Println("OpenFileDialog - ", sel)
	} else {
		fmt.Println("OpenFileDialog - No file selected")
	}

	sel, ok = tinyfiledialogs.OpenFileDialog("Source Files", homedir+"/", []string{"*.go", "*.mod"}, "Go files", true)
	if ok {
		fmt.Println("OpenFileDialog Multiple - ", sel)
	} else {
		fmt.Println("OpenFileDialog Multiple - No files selected")
	}

	sel, ok = tinyfiledialogs.SelectFolderDialog("Source Files", homedir)
	if ok {
		fmt.Println("SelectFolderDialog - ", sel)
	} else {
		fmt.Println("SelectFolderDialog - No folder selected")
	}

	sel, ok = tinyfiledialogs.SaveFileDialog("Source Files", homedir+"/", []string{"*.go", "*.mod"}, "Go files")
	if ok {
		fmt.Println("SaveFileDialog - ", sel)
	} else {
		fmt.Println("SaveFileDialog - Action canceled")
	}
}

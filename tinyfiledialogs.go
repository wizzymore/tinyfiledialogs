package tinyfiledialogs

/*
#cgo CFLAGS: -std=c89 -g3 -Wall -Wextra -Wdouble-promotion -Wconversion -Wno-sign-conversion -Wno-unused-parameter -Wno-unused-function -Wno-deprecated
#include "tinyfiledialogs.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// OpenFileDialog opens a file selection dialog.
//
// title - the title of the dialog.
//
// defaultPathOrFile - the default path or file to open. For directories, use a trailing slash.
//
// filePatterns is a list of file patterns to filter the files.
//
// singleFilterDescription is the description of the filter.
//
// allowMultipleSelection is a boolean to allow multiple file selection.
//
// Returns the selected file(s). If multiple files are selected, they are separated by a pipe character "|".
func OpenFileDialog(title string, defaultPathOrFile string, filePatterns []string, singleFilterDescription string, allowMultipleSelection bool) (path string, ok bool) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cDefaultPathOrFile := C.CString(defaultPathOrFile)
	defer C.free(unsafe.Pointer(cDefaultPathOrFile))

	var cFilePatterns **C.char
	{
		aFilePatterns := make([]*C.char, len(filePatterns))
		for i, pattern := range filePatterns {
			cstring := C.CString(pattern)
			aFilePatterns[i] = cstring
			defer C.free(unsafe.Pointer(cstring))
		}
		cFilePatterns = &aFilePatterns[0]
	}

	aSingleFilterDescription := C.CString(singleFilterDescription)
	defer C.free(unsafe.Pointer(aSingleFilterDescription))

	aAllowMultipleSelection := C.int(0)
	if allowMultipleSelection {
		aAllowMultipleSelection = C.int(1)
	}

	result := C.tinyfd_openFileDialog(cTitle, cDefaultPathOrFile, C.int(len(filePatterns)), cFilePatterns, aSingleFilterDescription, aAllowMultipleSelection)

	if result != nil {
		path = C.GoString(result)
	}

	return path, result != nil
}

// SaveFileDialog opens a file selection dialog.
//
// title - the title of the dialog.
//
// defaultPathOrFile - the default path or file to open. For directories, use a trailing slash.
//
// filePatterns is a list of file patterns to filter the files.
//
// singleFilterDescription is the description of the filter.
//
// Returns the selected file(s). If multiple files are selected, they are separated by a pipe character "|".
func SaveFileDialog(title string, defaultPathOrFile string, filePatterns []string, singleFilterDescription string) (path string, ok bool) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cDefaultPathOrFile := C.CString(defaultPathOrFile)
	defer C.free(unsafe.Pointer(cDefaultPathOrFile))

	var cFilePatterns **C.char
	{
		aFilePatterns := make([]*C.char, len(filePatterns))
		for i, pattern := range filePatterns {
			cstring := C.CString(pattern)
			aFilePatterns[i] = cstring
			defer C.free(unsafe.Pointer(cstring))
		}
		cFilePatterns = &aFilePatterns[0]
	}

	aSingleFilterDescription := C.CString(singleFilterDescription)
	defer C.free(unsafe.Pointer(aSingleFilterDescription))

	result := C.tinyfd_saveFileDialog(cTitle, cDefaultPathOrFile, (C.int)(len(filePatterns)), cFilePatterns, aSingleFilterDescription)

	if result != nil {
		path = C.GoString(result)
	}

	return path, result != nil
}

// title - the title of the dialog.
// defaultPath - the default path to open.
// The selected folder is not returned with a trailing slash compared to the tinyfd default behaviour.
// If you want the trailing slash use SelectFolderDialogT
func SelectFolderDialog(title string, defaultPath string) (path string, ok bool) {
	aTitle := C.CString(title)
	defer C.free(unsafe.Pointer(aTitle))

	aDefaultPath := C.CString(defaultPath)
	defer C.free(unsafe.Pointer(aDefaultPath))

	result := C.tinyfd_selectFolderDialog(aTitle, aDefaultPath)

	if result != nil {
		path = C.GoString(result)
	}

	return path, result != nil
}

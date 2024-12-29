package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	modOle32                 = syscall.NewLazyDLL("ole32.dll")
	modShell32               = syscall.NewLazyDLL("shell32.dll")
	modComdlg32              = syscall.NewLazyDLL("comdlg32.dll")
	procCoInitializeEx       = modOle32.NewProc("CoInitializeEx")
	procSHBrowseForFolderW   = modShell32.NewProc("SHBrowseForFolderW")
	procSHGetPathFromIDListW = modShell32.NewProc("SHGetPathFromIDListW")
)

const (
	COINIT_APARTMENTTHREADED = 0x2
	BIF_RETURNONLYFSDIRS     = 0x0001
	BIF_NEWDIALOGSTYLE       = 0x0040
)

type browseInfo struct {
	hwndOwner      syscall.Handle
	pidlRoot       uintptr
	pszDisplayName *uint16
	lpszTitle      *uint16
	ulFlags        uint32
	lpfn           uintptr
	lParam         uintptr
	iImage         int32
}

func initCOM() error {
	hr, _, _ := procCoInitializeEx.Call(0, COINIT_APARTMENTTHREADED)
	if hr != 0 {
		return fmt.Errorf("failed to initialize COM, HRESULT: 0x%x", hr)
	}
	return nil
}

func selectDirectory(title string) (string, error) {
	// Convert the title to a UTF-16 pointer
	titlePtr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return "", err
	}

	// Buffer for the selected path
	var buffer [syscall.MAX_PATH]uint16

	bi := browseInfo{
		hwndOwner:      0,
		pidlRoot:       0,
		pszDisplayName: &buffer[0],
		lpszTitle:      titlePtr,
		ulFlags:        BIF_RETURNONLYFSDIRS | BIF_NEWDIALOGSTYLE,
		lpfn:           0,
		lParam:         0,
		iImage:         0,
	}

	// Display the folder selection dialog
	ret, _, _ := procSHBrowseForFolderW.Call(uintptr(unsafe.Pointer(&bi)))
	if ret == 0 {
		return "", fmt.Errorf("folder selection canceled")
	}

	// Retrieve the path from the returned PIDL
	ok, _, _ := procSHGetPathFromIDListW.Call(ret, uintptr(unsafe.Pointer(&buffer[0])))
	if ok == 0 {
		return "", fmt.Errorf("failed to get selected folder path")
	}

	// Convert the path from UTF-16 to a Go string
	return syscall.UTF16ToString(buffer[:]), nil
}

func main() {
	// Initialize COM
	err := initCOM()
	if err != nil {
		fmt.Println("Error initializing COM:", err)
		os.Exit(1)
	}

	// Open folder selector
	selectedDirectory, err := selectDirectory("Select Document Root")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Println(selectedDirectory)
	}
}

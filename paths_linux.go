package gowatchprog

import (
	"os"
	"path/filepath"
)

var InstallPathRoot = map[ProgramContext]string{
	SystemService: filepath.Join("/", "usr", "local"),
	AllUsers:      filepath.Join("/", "usr", "local"),
	CurrentUser:   filepath.Join(os.Getenv("HOME"), ".local"),
}

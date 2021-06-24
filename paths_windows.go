package gowatchprog

import (
	"os"
)

var InstallPathRoot = map[ProgramContext]string{
	SystemService: os.Getenv("ALLUSERSPROFILE"),
	AllUsers:      os.Getenv("ALLUSERSPROFILE"),
	CurrentUser:   os.Getenv("LOCALAPPDATA"),
}

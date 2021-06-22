package gowatchprog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Get the path to the installation directory
func (p *Program) InstallDirectory(create bool) (string, error) {

	switch p.InstallContext {

	case AllUsers:
		return getEnvDirectory("ALLUSERSPROFILE", p.safeName(), create)

	case CurrentUser:
		return getEnvDirectory("LOCALAPPDATA", p.safeName(), create)

	default:
		return "", errors.New("invalid installation context")
	}
}

// Get path to the app data directory
func (p *Program) DataDirectory(create bool) (string, error) {

	switch p.StartupContext {

	case AllUsers:
		return getEnvDirectory("ALLUSERSPROFILE", p.safeName(), create)

	case CurrentUser:
		return getEnvDirectory("LOCALAPPDATA", p.safeName(), create)

	default:
		return "", errors.New("invalid installation context")
	}
}

// Create a directory path from an env variable and a final directory name
func getEnvDirectory(envVar, dirName string, create bool) (string, error) {

	// Get path to directory
	dirPath := os.Getenv(envVar)
	if len(dirPath) == 0 {
		return "", errors.New("unable to find environment path")
	}
	finalPath := filepath.Join(dirPath, dirName)

	// Create if create is true
	if create {

		if mkdirErr := os.MkdirAll(finalPath, 0644); mkdirErr != nil {
			return finalPath, fmt.Errorf("unable to create directory: %v. %v", finalPath, mkdirErr)
		}
	}

	return finalPath, nil
}

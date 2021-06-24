/**
file: paths.go

This file provides functions for getting paths to various
files and directories regarding the program and its contexts.

*/
package gowatchprog

import (
	"errors"
	"path/filepath"
	"regexp"
)

var ErrPathNotFound = errors.New("path not found")

// Get the path to the installation directory
func (p *Program) InstallDirectory(create bool) (string, error) {
	return getInstallDirectory(p.InstallContext, create)
}

// Get the path to the installation data directory
func (p *Program) DataDirectory(create bool) (string, error) {
	return getInstallDirectory(p.StartupContext, create)
}

// Get the path to the ExeFile inside the installation directory
func (p *Program) InstalledBinary() (string, error) {
	dir, err := p.InstallDirectory(false)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, p.ExeFile), nil
}

// Get the path to the install binary including cli arguments
func (p *Program) InstalledCommandLine() (string, error) {
	exePath, pathErr := p.InstalledBinary()
	if pathErr != nil {
		return "", pathErr
	}
	args := p.Args
	if args == nil {
		args = []string{}
	}
	return getCommandLine(exePath, args), nil
}

// Returns a name safe to use for a directory and registry key
func (p *Program) safeName() string {
	nonAscii := regexp.MustCompile(`(?i)[^A-Z0-9]`)
	return nonAscii.ReplaceAllLiteralString(p.Name, "-")
}

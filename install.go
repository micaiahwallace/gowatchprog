package gowatchprog

import (
	"os"
	"path"
)

// Install a service to the system
func (p *Program) Install(sourceDir string) error {

	// Test if source binary exists
	srcBin := path.Join(sourceDir, p.ExeFile)

	// Ensure install directory exists
	dstDir, dsterr := p.installDirectory()
	if dsterr != nil {
		return dsterr
	}
	if mkdsterr := os.MkdirAll(dstDir, 0644); mkdsterr != nil {
		return mkdsterr
	}
	dstBin := path.Join(dstDir, p.ExeFile)

	// Copy source binary into install dir
	return copyFileContents(srcBin, dstBin)
}

// Uninstall service from the system
func (p *Program) Uninstall() error {

	// Unregister startup
	if rmerr := p.DeregisterStartup(); rmerr != nil {
		return rmerr
	}

	// Get install directory
	dir, err := p.installDirectory()
	if err != nil {
		return err
	}

	// Remove installation directory
	return os.RemoveAll(dir)
}

// Check if the program is installed
func (p *Program) Installed() bool {

	// Get binary file path
	binpath, berr := p.installPathBin()
	if berr != nil {
		return false
	}

	// Test if installed binary exists
	if _, err := os.Stat(binpath); os.IsNotExist(err) {
		return false
	}

	return true
}

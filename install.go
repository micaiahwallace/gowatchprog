package gowatchprog

import (
	"fmt"
	"os"
	"path/filepath"
)

// Install ExeFile specified by sourceBin to the system
func (p *Program) Install(sourceBin string) error {

	// Ensure install directory exists
	dstDir, dsterr := p.InstallDirectory(true)
	if dsterr != nil {
		return dsterr
	}
	dstBin := filepath.Join(dstDir, p.ExeFile)

	// Copy source binary into install dir
	return copyFileContents(sourceBin, dstBin)
}

// Uninstall service from the system
func (p *Program) Uninstall() error {

	// Unregister startup
	if rmerr := p.RemoveStartup(); rmerr != nil {
		return fmt.Errorf("unable to remove startup: %v", rmerr)
	}

	// Get install directory
	dir, err := p.InstallDirectory(false)
	if err != nil {
		return err
	}

	// Remove installation directory
	if rderr := os.RemoveAll(dir); rderr != nil {
		return fmt.Errorf("unable to remove install dir: %v", rderr)
	}
	return nil
}

// Check if the program is installed
func (p *Program) Installed() bool {

	// Get binary file path
	binpath, berr := p.InstallPathBin()
	if berr != nil {
		return false
	}

	// Test if installed binary exists
	if _, err := os.Stat(binpath); os.IsNotExist(err) {
		return false
	}

	return true
}

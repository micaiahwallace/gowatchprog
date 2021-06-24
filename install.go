package gowatchprog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Install sourceBin binary file to p.ExeFile in the install directory
func (p *Program) InstallBinary(sourceBin string) error {

	// Ensure install directory exists
	dstDir, dsterr := p.InstallDirectory(true)
	if dsterr != nil {
		return dsterr
	}
	dstBin := filepath.Join(dstDir, p.ExeFile)

	// Copy source binary into install dir
	return copyFileContents(sourceBin, dstBin)
}

// Remove installation directories specified by InstallContext and StartupContext
func (p *Program) RemoveInstallation() error {

	erroredCmds := []string{}

	// Remove InstallContext install directory
	if riErr := removeContextDirectory(p.InstallContext); riErr != nil {
		erroredCmds = append(erroredCmds, riErr.Error())
	}

	// Remove StartupContext install directory (if different from InstallContext)
	if p.StartupContext != p.InstallContext {
		if rsErr := removeContextDirectory(p.StartupContext); rsErr != nil {
			erroredCmds = append(erroredCmds, rsErr.Error())
		}
	}

	// Return any removal errors
	if len(erroredCmds) > 0 {
		return fmt.Errorf("failed to remove install directories: %s", strings.Join(erroredCmds, ", "))
	}

	return nil
}

// Uninstall is equivalent to running RemoveStartup() then RemoveInstallation()
func (p *Program) Uninstall() error {

	erroredCmds := []string{}

	// Unregister startup
	if rsErr := p.RemoveStartup(); rsErr != nil {
		erroredCmds = append(erroredCmds, fmt.Sprintf("unable to remove startup: %v", rsErr))
	}

	// Remove installation directories
	if rdErr := p.RemoveInstallation(); rdErr != nil {
		erroredCmds = append(erroredCmds, rdErr.Error())
	}

	// Return any uninstall errors
	if len(erroredCmds) > 0 {
		return fmt.Errorf("uninstall failed: %s", strings.Join(erroredCmds, ", "))
	}

	return nil
}

// Check if the program is installed
func (p *Program) Installed() bool {

	// Get binary file path
	binpath, berr := p.InstalledBinary()
	if berr != nil {
		return false
	}

	// Test if installed binary exists
	if _, err := os.Stat(binpath); os.IsNotExist(err) {
		return false
	}

	return true
}

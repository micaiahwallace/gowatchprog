package gowatchprog

import (
	"errors"
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// Register the installed service startup in windows registry based on p.Context
//
// For ProgramContext AllUsers, prog.UserInstaller must be set to
// the full path to the user context installation command to be run
// when each user logs in for the first time after AllUsers install is run
func (p *Program) RegisterStartup() error {

	// Get current path to executable and args
	exePath, patherr := p.InstallPathBinWithArgs()
	if patherr != nil {
		return patherr
	}

	// Perform context specific startup logic
	switch p.StartupContext {

	case AllUsers:

		// Check for installer path
		if p.UserInstaller == "" {
			return errors.New("program.UserInstaller must be set before registering startup with ProgramContext AllUsers")
		}

		// Create active setup registry key
		if nkErr := createRegistryKey(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, p.safeName())); nkErr != nil {
			return nkErr
		}

		// Set installation version
		if addVerErr := writeRegistry(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, p.safeName()), "Version", "1"); addVerErr != nil {
			return addVerErr
		}

		// Set installation command for each user
		return writeRegistry(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, p.safeName()), "StubPath", p.UserInstaller)

	case CurrentUser:

		// Write the run registry key for the current user
		return writeRegistry(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, p.safeName(), exePath)

	default:
		return errors.New("unknown ProgramContext specified")
	}

	return nil
}

// Remove the installed service from startup
func (p *Program) RemoveStartup() error {
	switch p.StartupContext {
	case AllUsers:

		// remove active setup key, @todo find a way to remove startup key from all users
		if rmKeyErr := deleteRegistryKey(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, p.safeName())); rmKeyErr != nil {
			return rmKeyErr
		}
	case CurrentUser:

		// Remove the current user run key
		return removeRegistry(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, p.safeName())
	}

	return nil
}

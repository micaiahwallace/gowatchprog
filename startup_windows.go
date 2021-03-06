/**
Steps to register startup for AllUsers on windows:
		- Create an active setup registry key for the program
		- Add a Version value to the key, or bump it up if it exists
		- Add a StubPath value providing the user installation command

		The next time any user logs in, they will run the command if
		their hkey current user version for the program doesn't match.
		This command should be the path to an executable that runs the
		register startup with a CurrentUser StartupContext to complete the
		installation.

Steps to register startup for CurrentUser on windows:
		Simply add the command string value for the program to the registry run key
*/
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

	binPath, pathErr := p.InstalledBinary()
	if pathErr != nil {
		return pathErr
	}

	cmdPath := GetCommandLine(binPath, p.Args)
	pName := p.SafeName()

	switch p.StartupContext {

	case AllUsers:
		if p.UserInstaller == "" {
			return errors.New("UserInstaller must be set before registering startup with AllUsers StartupContext")
		}
		if _, _, nkErr := registry.CreateKey(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, pName), registry.SET_VALUE); nkErr != nil {
			return nkErr
		}
		if addVerErr := bumpActiveSetupVersion(registry.LOCAL_MACHINE, pName); addVerErr != nil {
			return addVerErr
		}
		return writeRegistry(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, pName), "StubPath", p.UserInstaller)

	case CurrentUser:
		return writeRegistry(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, pName, cmdPath)

	case SystemService:
		return errors.New("system service startup not yet implemented on windows")

	default:
		return ErrInvalidContext
	}

	return nil
}

// Remove the installed service from startup on windows
func (p *Program) RemoveStartup() error {

	pName := p.SafeName()

	switch p.StartupContext {

	case AllUsers:
		// remove active setup key, @todo find a way to remove startup key from all users
		return deleteRegistryKey(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, pName))

	case CurrentUser:
		return removeRegistry(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, pName)

	case SystemService:
		return errors.New("system service startup not yet implemented on windows")

	default:
		return ErrInvalidContext
	}

	return nil
}

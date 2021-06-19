package gowatchprog

import "golang.org/x/sys/windows/registry"

// Register the installed service startup in windows registry
func (p *Program) RegisterStartup() error {

	// Get current path to executable and args
	exePath, patherr := p.installPathBinWithArgs()
	if patherr != nil {
		return patherr
	}

	// Perform context specific startup logic
	switch p.Context {
	case AllUsers:
		return writeRegistry(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, p.safeName(), exePath)
	default:
	}

	return nil
}

// Remove the installed service from startup
func (p *Program) RemoveStartup() error {
	switch p.Context {
	case AllUsers:
		return removeRegistry(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, p.safeName())
	}

	return nil
}

package gowatchprog

// Register the installed service startup in windows registry
func (p *Program) RegisterStartup() error {

	switch p.Context {
	case AllUsers:
		return p.registerStartupAllUsers()
	default:
	}
}

func (p *Program) registerStartupAllUsers() error {

	// Open registry key
	key, kerr := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE|registry.QUERY_VALUE)
	if kerr != nil {
		return kerr
	}

	// Get current path to executable and args
	exePath, patherr := p.installPathBinWithArgs()
	if patherr != nil {
		return patherr
	}

	// Add the string value
	if kverr := key.SetStringValue(p.safeName(), exePath); kverr != nil {
		return kverr
	}

	return nil
}

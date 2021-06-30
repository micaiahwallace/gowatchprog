package gowatchprog

import "errors"

// Register the installed service startup on linux
func (p *Program) RegisterStartup() error {
	switch p.StartupContext {

	case AllUsers:
		return errors.New("all users startup not implemented on linux")

	case CurrentUser:
		return errors.New("current user startup not implemented on linux")

	case SystemService:
		return errors.New("system service startup not implemented on linux")

	default:
		return ErrInvalidContext
	}
}

// Deregister the installed service from startup on linux
func (p *Program) RemoveStartup() error {
	switch p.StartupContext {

	case AllUsers:
		return errors.New("remove startup for allusers not implemented on linux")

	case CurrentUser:
		return errors.New("remove startup for currentuser not implemented on linux")

	case SystemService:
		return errors.New("remove startup for systemservice not implemented on linux")

	default:
		return ErrInvalidContext
	}
}

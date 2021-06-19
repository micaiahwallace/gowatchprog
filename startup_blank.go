// +build !windows

package gowatchprog

import "errors"

// Register the installed service startup
func (p *Program) RegisterStartup() error {
	return errors.New("register startup not implemented on this os")
}

// Deregister the installed service from startup
func (p *Program) RemoveStartup() error {
	return errors.New("remove startup not implemented on this os")
}

// +build !windows

package gowatchprog

import "errors"

// Register the installed service startup
func (p *Program) RegisterStartup() error {
	return errors.New("register not implemented on this os")
}

// Deregister the installed service from startup
func (p *Program) DeregisterStartup() error {
	return errors.New("deregister not implemented on this os")
}

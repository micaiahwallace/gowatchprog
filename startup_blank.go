// +build !windows

package gowatchprog

import "errors"

// Register the installed service startup
func (p *Program) RegisterStartup() error {
	return errors.New("register not implemented on this os")
}

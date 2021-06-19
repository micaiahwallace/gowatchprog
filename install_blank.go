// +build !windows

package gowatchprog

import "errors"

// Get the path to the installation directory
func (p *Program) installDirectory() (string, error) {
	return "", errors.New("install not implemented on this o")
}

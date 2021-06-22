// +build !windows

package gowatchprog

import "errors"

// Get the path to the installation directory
func (p *Program) InstallDirectory(create bool) (string, error) {
	return "", errors.New("installdirectory not implemented on this os")
}

// Get the path to the app data directory
func (p *Program) DataDirectory(create bool) (string, error) {
	return "", errors.New("datadirectory not implemented on this os")
}

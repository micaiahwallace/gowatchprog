package gowatchprog

import (
	"errors"
	"os"
	"path"
)

// Get the path to the installation directory
func (p *Program) installDirectory() (string, error) {

	switch p.Context {
	case AllUsers:

		// Check for all users profile env variable
		profilePath := os.Getenv("ALLUSERSPROFILE")
		if len(profilePath) == 0 {
			return "", errors.New("installation path not found")
		}
		return path.Join(profilePath, p.safeName()), nil

	default:
		return "", errors.New("invalid installation context")
	}
}

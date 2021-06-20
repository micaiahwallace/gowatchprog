package gowatchprog

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

// Get the path to the install binary
func (p *Program) installPathBin() (string, error) {

	// Merge install directory with exe file name
	dir, err := p.installDirectory()
	if err != nil {
		return "", err
	}
	return path.Join(dir, p.ExeFile), nil
}

// Get the path to the install binary including cli arguments
func (p *Program) installPathBinWithArgs() (string, error) {

	// Get current path to executable and args
	exePath, patherr := p.installPathBin()
	if patherr != nil {
		return "", patherr
	}

	// Append program arguments to executable string
	exePath = fmt.Sprintf(`"%s"`, exePath)
	if p.Args != nil {
		parts := append([]string{exePath}, p.Args...)
		exePath = strings.Join(parts, " ")
	}

	return exePath, nil
}

// Returns a name safe to use for a directory and registry key
func (p *Program) safeName() string {
	nonAscii := regexp.MustCompile(`(?i)[^A-Z0-9]`)
	return nonAscii.ReplaceAllLiteralString(p.Name, "-")
}

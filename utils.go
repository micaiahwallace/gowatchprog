package gowatchprog

import (
	"fmt"
	"strings"
)

// Wrap the binary path in quotes and append the space
// separated args at the end, return the final string
func GetCommandLine(binPath string, args []string) string {

	// ensure args are not nil
	var safeArgs = args
	if safeArgs == nil {
		safeArgs = []string{}
	}

	// add quotes around all paths
	quotedPath := fmt.Sprintf(`"%s"`, binPath)

	// place bin and args together then join with a space separator
	parts := append([]string{quotedPath}, safeArgs...)
	return strings.Join(parts, " ")
}

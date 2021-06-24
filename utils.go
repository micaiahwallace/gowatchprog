package gowatchprog

import (
	"fmt"
	"strings"
)

// Create a runnable command line string with arguments appended
func getCommandLine(binPath string, args []string) string {
	quotedPath := fmt.Sprintf(`"%s"`, binPath)
	parts := append([]string{quotedPath}, args...)
	return strings.Join(parts, " ")
}

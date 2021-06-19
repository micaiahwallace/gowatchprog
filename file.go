package gowatchprog

import (
	"io"
	"os"
)

// Copy a source file to a destination
func copyFileContents(src, dst string) (err error) {

	// Open source file handle
	in, err := os.Open(src)
	if err != nil {
		return
	}

	// Close source file handle later
	defer in.Close()

	// Create and open destination file handle
	out, err := os.Create(dst)
	if err != nil {
		return
	}

	// Close destination handle later
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Copy file contents from src into dst file
	if _, err = io.Copy(out, in); err != nil {
		return
	}

	// Do final write commit before returning
	err = out.Sync()
	return
}

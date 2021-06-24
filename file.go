package gowatchprog

import (
	"fmt"
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

// Get an installation directory based on a ProgramContext
func getInstallDirectory(ctx ProgramContext, create bool) (string, error) {
	installRoot := InstallPathRoot[ctx]
	if installRoot == "" {
		return "", ErrPathNotFound
	}
	if create {
		if err := makeDirectory(installRoot); err != nil {
			return "", err
		}
	}
	return installRoot, nil
}

// Create the full path to the directory if it doesn't exist
func makeDirectory(dir string) error {
	if mkdirErr := os.MkdirAll(dir, 0644); mkdirErr != nil {
		return fmt.Errorf("create failed: %v. (%s)", dir, mkdirErr.Error())
	}
	return nil
}

// Remove a directory specified by a program context
func removeContextDirectory(ctx ProgramContext) error {
	dir, err := getInstallDirectory(ctx, false)
	if err != nil {
		return err
	}
	if rdErr := os.RemoveAll(dir); rdErr != nil {
		return fmt.Errorf("unlink failed: %s (%s)", dir, rdErr.Error())
	}
	return nil
}

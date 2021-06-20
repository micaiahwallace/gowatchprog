package gowatchprog

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// Write a string value to a specified registry key on windows
func writeRegistry(key registry.Key, path, keyName, keyValue string) error {

	// Open registry key
	key, kerr := registry.OpenKey(key, path, registry.SET_VALUE)
	if kerr != nil {
		return fmt.Errorf("unable to open registry: %v. %v", path, kerr)
	}

	// Add the string value
	if kverr := key.SetStringValue(keyName, keyValue); kverr != nil {
		return fmt.Errorf("unable to set registry value: %v. %v", keyName, kverr)
	}

	return nil
}

// Remove a value from a specified registry path
func removeRegistry(key registry.Key, path, keyName string) error {

	// Open registry key
	key, kerr := registry.OpenKey(key, path, registry.SET_VALUE)
	if kerr != nil {
		return fmt.Errorf("unable to open registry: %v. %v", path, kerr)
	}

	// Remove the specified value
	if err := key.DeleteValue(keyName); err != nil {
		return fmt.Errorf("unable to delete registry value: %v. %v", keyName, err)
	}

	return nil
}

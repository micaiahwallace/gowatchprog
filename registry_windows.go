package gowatchprog

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// Create a new registry key
func createRegistryKey(key registry.Key, path string) error {

	// Create new key at specified hive and path
	if _, _, err := registry.CreateKey(key, path, registry.SET_VALUE); err != nil {
		return fmt.Errorf("unable to create registry key: %v. %v", path, err)
	}

	return nil
}

// Delete a registry key
func deleteRegistryKey(key registry.Key, path string) error {

	// Create new key at specified hive and path
	if err := registry.DeleteKey(key, path); err != nil {
		return fmt.Errorf("unable to delete registry key: %v. %v", path, err)
	}

	return nil
}

// Write a string value to a specified registry key on windows
func writeRegistry(key registry.Key, path, keyName, keyValue string) error {

	// Open registry key
	regKey, kerr := registry.OpenKey(key, path, registry.SET_VALUE)
	if kerr != nil {
		return fmt.Errorf("unable to open registry: %v. %v", path, kerr)
	}

	// Add the string value
	if kverr := regKey.SetStringValue(keyName, keyValue); kverr != nil {
		return fmt.Errorf("unable to set registry value: %v. %v", keyName, kverr)
	}

	return nil
}

// Remove a value from a specified registry path
func removeRegistry(key registry.Key, path, keyName string) error {

	// Open registry key
	regKey, kerr := registry.OpenKey(key, path, registry.SET_VALUE)
	if kerr != nil {
		return fmt.Errorf("unable to open registry: %v. %v", path, kerr)
	}

	// Remove the specified value
	if err := regKey.DeleteValue(keyName); err != nil {
		return fmt.Errorf("unable to delete registry value: %v. %v", keyName, err)
	}

	return nil
}

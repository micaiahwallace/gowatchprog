package gowatchprog

import (
	"fmt"
	"strconv"

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

// Create or bump up the version number for an active setup registry key
func bumpActiveSetupVersion(key registry.Key, name string) error {

	// Get the registry key path
	keyPath := fmt.Sprintf(`SOFTWARE\Microsoft\Active Setup\Installed Components\%v`, name)

	// Open registry key
	regKey, kerr := registry.OpenKey(key, keyPath, registry.SET_VALUE|registry.QUERY_VALUE)
	if kerr != nil {
		return fmt.Errorf("unable to open registry: %v. %v", keyPath, kerr)
	}

	// Get the current version number
	currentVer, _, verr := regKey.GetStringValue("Version")
	if verr == registry.ErrNotExist {
		currentVer = "0"
	} else if verr != nil {
		return verr
	}

	// Convert version to an int
	cvNum, convErr := strconv.Atoi(currentVer)
	if convErr != nil {
		return convErr
	}

	// Get the next version in string form
	newVersion := strconv.Itoa(cvNum + 1)

	// Write the new version to the registry
	if wrErr := regKey.SetStringValue("Version", newVersion); wrErr != nil {
		return wrErr
	}

	return nil
}

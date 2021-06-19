package gowatchprog

import "golang.org/x/sys/windows/registry"

// Write a string value to a specified registry key on windows
func writeRegistry(key registry.Key, path, keyName, keyValue string) error {

	// Open registry key
	key, kerr := registry.OpenKey(key, path, registry.SET_VALUE|registry.QUERY_VALUE)
	if kerr != nil {
		return kerr
	}

	// Add the string value
	if kverr := key.SetStringValue(keyName, keyValue); kverr != nil {
		return kverr
	}

	return nil
}

// Remove a value from a specified registry path
func removeRegistry(key registry.Key, path, keyName string) error {

	// Open registry key
	key, kerr := registry.OpenKey(key, path, registry.SET_VALUE|registry.QUERY_VALUE)
	if kerr != nil {
		return kerr
	}

	// Remove the specified value
	return key.DeleteValue(keyName)
}

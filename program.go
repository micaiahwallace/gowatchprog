package gowatchprog

import "time"

type ProgramContext int

const (

	// The installation context is global for all users
	AllUsers ProgramContext = iota

	// The installation context is set for the current user
	CurrentUser
)

// Definition for a service to manage
type Program struct {

	// Name of the service
	Name string

	// Filename of service binary
	ExeFile string

	// Arguments to append when service is run
	Args []string

	// Path to installer for user context when AllUsers context is specified
	UserInstaller string

	// Installation context of the service
	InstallContext ProgramContext

	// Startup context of the service
	StartupContext ProgramContext

	// Watchdog retry count before failing, -1 for unlimited
	WatchRetries int

	// Watchdog interval between retries
	WatchRetryWait time.Duration

	// Watchdog factor to increase wait interval each failed attempt
	WatchRetryIncrease int
}

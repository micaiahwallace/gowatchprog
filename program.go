package gowatchprog

import "time"

// Specifies the type of program installation, so that Program can customize how it installs and configures startup on each system.
type ProgramContext int

const (

	// A global context to install and run user session software for all users
	AllUsers ProgramContext = iota

	// A local context to install and run user session software for the current user
	CurrentUser

	// A global context to install and run a system level service not running in a user context
	SystemService
)

// Program defines a program installation that can be used to install, uninstall or start a watchdog based on the defined properties.
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

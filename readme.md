# GoWatchProg

This module was created to help manage background services on machines by providing the following core feature sets.

- Install to a proper location based on installation context (user, system, allusers)
- Register service autostart based on service run context (user, system, allusers)
- Watchdog autorestart (retry count, retry delay, retry delay increase)
- Install updates remotely

## Install

```shell
go get github.com/micaiahwallace/gowatchprog
```

## Module Definition

```go
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
```

## Core Exported Functions

```go
// Start the watchdog runner
func (p *Program) RunWatchdog(errs chan error, msgs chan string, quit chan interface{}) 

// Install ExeFile specified by sourceBin to the system
func (p *Program) Install(sourceBin string) error

// // Register the installed service startup in windows registry based on p.Context
func (p *Program) RegisterStartup() error

// Remove the service from automatic startup
func (p *Program) RemoveStartup() error

// Uninstall ExeFile from the system
func (p *Program) Uninstall() error
```

## Exported Helper Functions

```go
// Check if the program is installed
func (p *Program) Installed() bool

// Get the path to the installation directory
func (p *Program) InstallDirectory(create bool) (string, error) 

// Get path to the app data directory
func (p *Program) DataDirectory(create bool) (string, error)

// Get the path to the install binary
func (p *Program) InstallPathBin() (string, error)

// Get the path to the install binary including cli arguments
func (p *Program) InstallPathBinWithArgs() (string, error)
```

## Constants

```go
const (

	// The app should run under all users interactively
	AllUsers ProgramContext = 0

	// The installation context is set for the current user
	CurrentUser ProgramContext = 1
)
```

## Example 

You can see gowatchprog in action on one of my other projects [GoScreenMonit](https://github.com/micaiahwallace/goscreenmonit)

## Roadmap

- [x] Windows support (Full context set not implemented for install / startup)
- [ ] Remote update feature
- [ ] Mac OS support
- [ ] Linux support
- [ ] Features accessible from command line utility (for non-go services)

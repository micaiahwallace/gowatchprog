# WatchProg - Service manager

```shell
go get github.com/micaiahwallace/gowatchprog
```

This module was created to help manage background services on machines by providing the following core feature sets.

- Install to a proper location
- Register service autostart
- Watchdog autorestart
- Install updates and restart service

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

	// Installation and execution context of the service
	Context ProgramContext

	// Watchdog retry count before failing, -1 for unlimited
	WatchRetries int

	// Watchdog interval between retries
	WatchRetryWait time.Duration

	// Watchdog factor to increase wait interval each failed attempt
	WatchRetryIncrease int
}
```

## Exported Functions

```go
// Start the watchdog runner
func (p *Program) RunWatchdog(quit chan int)

// Check if the program is installed
func (p *Program) Installed() bool

// Install ExeFile from sourceDir to the system
func (p *Program) Install(sourceDir string) error

// Register the service to startup automatically, optionally with watchdog service
func (p *Program) RegisterStartup(watchdog bool) error

// Remove the service from automatic startup
func (p *Program) RemoveStartup() error

// Uninstall ExeFile from the system
func (p *Program) Uninstall() error
```

## Constants

```go
const (

	// The app should run under all users interactively
	AllUsers ProgramContext = 0
)
```

## Example 

You can see gowatchprog in action on one of my other projects [GoScreenMonit](https://github.com/micaiahwallace/goscreenmonit)

## Roadmap

- [x] Windows support (AllUsers context support only for now)
- [ ] Remote update feature
- [ ] Mac OS support
- [ ] Linux support
- [ ] Features accessible from command line utility (for non-go services)
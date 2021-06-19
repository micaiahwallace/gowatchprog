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

	// Execution context of the service
	Context ProgramContext
}


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

## Example 

View the samples in the example directory on use cases.

## Roadmap

- [] Core system features accessible via module import [install, autostart, watchdog]
- [] Windows support
- [] Mac OS support
- [] Linux support
- [] Remote update feature
- [] Features accessible from command line utility (for non-go services)
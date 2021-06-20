package gowatchprog

import "time"

type ProgramContext int

const (

	// The app should run under all users interactively
	AllUsers ProgramContext = iota
)

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
	watchRetries int

	// Watchdog interval between retries
	watchRetryWait time.Duration

	// Watchdog factor to increase wait interval each failed attempt
	watchRetryIncrease int
}

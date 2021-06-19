package gowatchprog

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

	// Execution context of the service
	Context ProgramContext
}

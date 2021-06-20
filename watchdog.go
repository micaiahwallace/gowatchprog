package gowatchprog

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"

	"github.com/mitchellh/go-ps"
)

const PID_FILE = "watchdog.pid"

// Start the watchdog runner
func (p *Program) RunWatchdog(quit chan int) {

	// Register a new pid file
	if !p.createWatchdogLock() {
		log.Println("unable to create pid file lock")
		quit <- 0
		return
	}

	// Get exe path and args
	exePath, eErr := p.installPathBin()
	if eErr != nil {
		log.Println("unable to get installation binary path")
		quit <- 0
		return
	}

	// Start loop to execute service
	failCount := 0
	for {

		// Execute command with arguments
		cmd := exec.Command(exePath, p.Args...)
		log.Printf("watchdog starting service attempt: %d\n", failCount)
		runErr := cmd.Run()
		if runErr != nil {
			failCount++
		}
		log.Printf("service completed with error: %v\n", runErr)

		// Check if retries exceeded
		if failCount >= p.watchRetries {
			log.Println("retry count exceeded, now exiting")
			break
		}

		// Wait configured duration before retrying
		time.Sleep(p.watchRetryWait * time.Duration(p.watchRetryIncrease*failCount))
	}

	// Remove pid file after completing watchdog
	if !p.removeWatchdogLock() {
		log.Println("unable to remove pid file lock")
	}

	quit <- 0
}

// Get a lock to run the watchdog using pid file method
func (p *Program) createWatchdogLock() bool {

	// get pid file path
	installDir, derr := p.installDirectory()
	if derr != nil {
		return false
	}
	pidPath := path.Join(installDir, PID_FILE)

	// Test if can create lock
	if !p.canWatchdogLock(pidPath) {
		return false
	}

	// create pid file with current pid and return true
	curPid := os.Getpid()
	err := os.WriteFile(pidPath, []byte(strconv.Itoa(curPid)), 0644)
	return err == nil
}

// Remove a watchdog pid lock file
func (p *Program) removeWatchdogLock() bool {

	// get pid file path
	installDir, derr := p.installDirectory()
	if derr != nil {
		return false
	}
	pidPath := path.Join(installDir, PID_FILE)

	// Remove the file
	return os.Remove(pidPath) == nil
}

// returns true if watchdog is able to get a lock
func (p *Program) canWatchdogLock(pidPath string) bool {

	pidContents, ferr := os.ReadFile(pidPath)
	if ferr != nil {

		// Test if pid file exists
		if os.IsNotExist(ferr) {
			return true
		}
	}

	pidNum, nerr := strconv.Atoi(string(pidContents))
	if nerr != nil {
		return false
	}

	proc, perr := ps.FindProcess(pidNum)
	if perr != nil {
		return false
	}

	currentBinFile := path.Base(os.Args[0])

	// Test if actual process still exists at pid
	return proc.Executable() != currentBinFile
}

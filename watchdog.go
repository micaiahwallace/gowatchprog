package gowatchprog

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/go-ps"
)

const PID_FILE = "watchdog.pid"

// Start the watchdog runner
func (p *Program) RunWatchdog(errs chan error, msgs chan string, quit chan interface{}) {

	// Register a new pid file
	if cwerr := p.createWatchdogLock(); cwerr != nil {
		errs <- fmt.Errorf("unable to create pid file lock: %v", cwerr)
		close(quit)
		return
	}

	// Get exe path and args
	exePath, eErr := p.InstalledBinary()
	if eErr != nil {
		errs <- fmt.Errorf("unable to get installation binary path: %v", eErr)
		close(quit)
		return
	}

	// Start loop to execute service
	failCount := 0
	for {

		// Execute command with arguments
		msgs <- fmt.Sprintf("watchdog starting service attempt: %d\n", failCount)
		cmd := exec.Command(exePath, p.Args...)
		runErr := cmd.Run()
		if runErr != nil {
			failCount++
		}
		msgs <- fmt.Sprintf("service completed with error: %v\n", runErr)

		// Check if retries exceeded
		if failCount >= p.WatchRetries && p.WatchRetries != -1 {
			msgs <- "retry count exceeded, now exiting"
			break
		}

		// Wait configured duration before retrying
		time.Sleep(p.WatchRetryWait * time.Duration(p.WatchRetryIncrease*failCount))
	}

	// Remove pid file after completing watchdog
	if rmwErr := p.removeWatchdogLock(); rmwErr != nil {
		errs <- rmwErr
	}

	close(quit)
}

// Get a lock to run the watchdog using pid file method
func (p *Program) createWatchdogLock() error {

	// get pid file path
	dataDir, derr := p.DataDirectory(true)
	if derr != nil {
		return derr
	}
	pidPath := filepath.Join(dataDir, PID_FILE)

	// Test if can create lock
	if lerr := p.canWatchdogLock(pidPath); lerr != nil {
		return lerr
	}

	// create pid file with current pid and return true
	curPid := os.Getpid()
	return os.WriteFile(pidPath, []byte(strconv.Itoa(curPid)), 0644)
}

// Remove a watchdog pid lock file
func (p *Program) removeWatchdogLock() error {

	// get pid file path
	dataDir, derr := p.DataDirectory(false)
	if derr != nil {
		return derr
	}
	pidPath := filepath.Join(dataDir, PID_FILE)

	// Remove the file
	return os.Remove(pidPath)
}

// returns true if watchdog is able to get a lock
func (p *Program) canWatchdogLock(pidPath string) error {

	pidContents, ferr := os.ReadFile(pidPath)

	// If PID file cannot be read for some reason, don't allow lock creation
	if os.IsPermission(ferr) || os.IsTimeout(ferr) {
		return ferr
	}

	// If PID file doesn't exist, allow lock creation
	if os.IsNotExist(ferr) {
		return nil
	}

	pidNum, nerr := strconv.Atoi(strings.TrimSpace(string(pidContents)))

	// Unable to parse PID value from file, allow lock creation
	if nerr != nil {
		return nil
	}

	proc, _ := ps.FindProcess(pidNum)

	// Process was not found or there was an error getting the process, allow lock creation
	if proc == nil {
		return nil
	}

	// Ensure pid doesn't refer to same process binary
	if proc.Executable() != p.ExeFile {
		return nil
	}

	return errors.New("watchdog already running")
}

/**

reorganization brainstorm:

// Get the path to watchdog pid file
getPidPath
	return path(dataDir, PID_FILE)

// get a reference to the process identified in the pid file if it exists
getPidFileProcess
	read pid file
	convert pid string to int
	get process from pid
	return process

// validate and create a watchdog pid file
createWatchdogLock
	if !canLockWatchdog {
		return err
	}
	writePidFile(pidPath)

// validate if pid file can be created
canLockWatchdog(pid)
	if !fileExists {
		return true
	}
	if getPidFileProcess.Exe == currentExeFile {
		return false
	}
	return true
*/

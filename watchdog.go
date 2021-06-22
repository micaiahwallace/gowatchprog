package gowatchprog

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
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
	exePath, eErr := p.InstallPathBin()
	if eErr != nil {
		errs <- fmt.Errorf("unable to get installation binary path: %v", eErr)
		close(quit)
		return
	}

	// Start loop to execute service
	failCount := 0
	for {

		// Execute command with arguments
		cmd := exec.Command(exePath, p.Args...)
		msgs <- fmt.Sprintf("watchdog starting service attempt: %d\n", failCount)
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
	if !p.removeWatchdogLock() {
		errs <- errors.New("unable to remove pid file lock")
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
func (p *Program) removeWatchdogLock() bool {

	// get pid file path
	dataDir, derr := p.DataDirectory(false)
	if derr != nil {
		return false
	}
	pidPath := filepath.Join(dataDir, PID_FILE)

	// Remove the file
	return os.Remove(pidPath) == nil
}

// returns true if watchdog is able to get a lock
func (p *Program) canWatchdogLock(pidPath string) error {

	pidContents, ferr := os.ReadFile(pidPath)
	if ferr != nil {

		// Test if pid file exists
		if os.IsNotExist(ferr) {
			return nil
		}
	}

	pidNum, nerr := strconv.Atoi(strings.TrimSpace(string(pidContents)))
	if nerr != nil {

		// Invalid pid format, allow lock creation
		return nil
	}

	proc, perr := ps.FindProcess(pidNum)
	if perr != nil || proc == nil {
		return perr
	}

	currentBinFile := path.Base(os.Args[0])

	// Ensure pid doesn't refer to same process binary
	if proc.Executable() == currentBinFile {
		return errors.New("watchdog already running")
	}

	return nil
}

/*
  UTILS Package
	helper functions for every case
*/

package utils

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

// Excuse for something went wrong
func Excuse() string {
	// http://pages.cs.wisc.edu/~ballard/bofh/excuses
	reasons := []string{
		"electromagnetic radiation from satellite debris",
		"global warming",
		"hardware stress fractures",
		"magnetic interference from money/credit cards",
		"sounds like a Windows problem, try calling Microsoft support",
		"temporary routing anomaly",
		"fat electrons in the lines",
		"floating point processor overflow",
		"POSIX compliance problem",
		"monitor resolution too high",
		"Decreasing electron flux",
		"first Saturday after first full moon in Winter",
		"CPU radiator broken",
		"positron router malfunction",
		"cellular telephone interference",
		"piezo-electric interference",
		"dynamic software linking table corrupted",
		"heavy gravity fluctuation, move computer to floor rapidly",
		"terrorist activities",
		"not enough memory, go get system upgrade",
	}
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	return reasons[rand.Intn(len(reasons))]
}

// GetCaller - get the name of the function which called me and fileline number
func GetCaller() (string, string) {
	fpcs := make([]uintptr, 1)
	// Skip 3 levels to get the caller
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "unknown", "unknown"
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		return "nil", "unknown"
	}
	// return the file name and line number
	file, line := caller.FileLine(fpcs[0] - 1)
	return caller.Name(), file + ":" + strconv.Itoa(line)
}

// FileExists - reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// GetDateTimestamp - returns nicely formated timestamp
func GetDateTimestamp() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
}

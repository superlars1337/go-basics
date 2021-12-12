package utils

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

// ExecuteCliCommand - wrapper : runs a bash command, output is stdout
func ExecuteCliCommand(cliCommand *exec.Cmd, silent bool) (err error) {
	_, _, err = executeSilentCLICommandWithResultInDirectory(cliCommand, "", silent)
	return err
}

// ----------------------------  Internal --------------------------------------------------------

// Run bash command
func executeSilentCLICommandWithResultInDirectory(cliCommand *exec.Cmd, directory string, silent bool) (stdout string, stderr string, error error) {
	cliCommand.Env = os.Environ()

	var outBuffer, errBuffer bytes.Buffer
	if silent {
		cliCommand.Stdout = &outBuffer
		cliCommand.Stderr = &errBuffer
	} else {
		cliCommand.Stdout = io.MultiWriter(&outBuffer, os.Stdout)
		cliCommand.Stderr = io.MultiWriter(&errBuffer, os.Stderr)
	}
	if directory != "" {
		cliCommand.Dir = directory
	}

	// PrintVerbose(cliCommand.String())
	if err := cliCommand.Run(); err != nil {
		return "", errBuffer.String(), err
	}
	return outBuffer.String(), errBuffer.String(), nil
}

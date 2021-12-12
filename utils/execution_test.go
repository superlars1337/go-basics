package utils

import (
	"os/exec"
)

func ExampleExecuteCliCommand() {
	cliCommand := exec.Command("echo", "hello", "world")
	_ = ExecuteCliCommand(cliCommand, false)
	// Output:
	//hello world
}

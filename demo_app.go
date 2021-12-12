package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/superlars1337/go-basics/log"
	"github.com/superlars1337/go-basics/utils"
)

func printMessages() {
	log.Message("A Message")
	log.Success("Success Message")
	log.Info("Info Message")
	log.Debug("Debug Message")
	log.Warn("Warn Message")
	err := errors.New("an error")
	log.Error(err, "Error Message")
	fmt.Println()
}

func main() {
	// utils
	name, _ := utils.ReadLine("Wie lautet ihr Name?")
	utils.PrintMessage("Hallo " + name)
	fmt.Println()

	utils.PrintInformation("Example CLI call")
	cliCommand := exec.Command("ls", "-l", "-a")
	err := utils.ExecuteCliCommand(cliCommand, false)
	log.CheckErrFatal(err)

	// logger
	debug := true
	utils.PrintSubTitle("Our pretty printer")
	log.Init(log.LOG_PRETTY, debug)
	printMessages()

	utils.PrintSubTitle("phuslu pretty printer")
	log.Init(log.PHUSLU_PRETTY, debug)
	printMessages()

	utils.PrintSubTitle("phuslu json printer")
	log.Init(log.PHUSLU_JSON, debug)
	printMessages()

	utils.PrintSubTitle("End Demo")
	log.Init(log.LOG_PRETTY, debug)
	// log.CheckErrFatal(errors.New("DEMO END"))
	log.ExitNow("Have a nice Day")
}

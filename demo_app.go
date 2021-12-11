package main

import (
	"errors"
	"fmt"

	"github.com/superlars1337/go-basics/log"
	"github.com/superlars1337/go-basics/utils"
)

func printMessages() {
	log.Info("Info Message")
	log.Debug("Debug Message")
	log.Warn("Warn Message")
	err := errors.New("an error")
	log.Error(err, "Error Message")
	fmt.Println()
}

func main() {
	utils.PrintTitle("Log Output Demonstration")

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

	utils.PrintSubTitle("End Demo with panic")
	log.Init(log.LOG_PRETTY, debug)
	log.CheckErrFatal(errors.New("DEMO END"))
}

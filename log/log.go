/*
  LOG Package
	for easy exchanging the used logger package
	eg:
	- https://github.com/phuslu/log         (*)
	- https://github.com/sirupsen/logrus
	OR
	- log to stdout (in color) .. for CLI applications
*/

package log

import (
	"os"

	phuslu "github.com/phuslu/log"
	"github.com/superlars1337/go-basics/utils"
)

// define logger Ids
const (
	LOG_PRETTY    = 0 // own pretty printer with emojis , DEFAULT LOGGER
	PHUSLU_PRETTY = 1
	PHUSLU_JSON   = 2
	// define more loggers here
)

// inital: debug is set to "false"
var (
	logDebug  bool = false
	logLogger int  = LOG_PRETTY
)

// Init - setup logger and mode (phuslu)
func Init(logger int, debug bool) {

	// global setting
	logLogger = logger
	logDebug = debug

	// phuslu logsetting
	phusluLogLevel := phuslu.InfoLevel
	if debug {
		phusluLogLevel = phuslu.DebugLevel
	}

	// init loggers
	switch logger {
	case PHUSLU_PRETTY:
		phuslu.DefaultLogger = phuslu.Logger{
			Level:      phusluLogLevel,
			TimeFormat: "2006-01-02 15:04:05",
			Writer: &phuslu.ConsoleWriter{
				ColorOutput:    phuslu.IsTerminal(os.Stderr.Fd()),
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	case PHUSLU_JSON:
		phuslu.DefaultLogger = phuslu.Logger{
			Level:      phusluLogLevel,
			TimeFormat: "2006-01-02 15:04:05",
		}
	}

	Debug("Debug Mode is on")
}

// Info - print infolog
func Info(s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Info().Msg(s)
	default:
		utils.PrintInformation("INFO: " + s)
	}
}

// Debug - print debuglog
func Debug(s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Debug().Msg(s)
	default:
		if logDebug {
			utils.PrintDebug("DBUG: " + s)
		}
	}
}

// Warn - print warnlog
func Warn(s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Warn().Msg(s)
	default:
		utils.PrintWarning("WARN: " + s)
	}
}

// Error - print errorlog
func Error(err error, s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Error().Err(err).Msg(s)
	default:
		utils.PrintError("EROR: " + s + "(" + err.Error() + ")")
	}
}

// ErrorS - print errorlog (string only)
func ErrorS(s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Error().Msg(s)
	default:
		utils.PrintError(s)
	}
}

// Fatal - print errorlog and exit program # panic
func Fatal(err error, s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Fatal().Err(err).Msg(s)
	default:
		utils.PrintError("FATL: " + s + " (" + err.Error() + ")")
	}
	PanicNow()
}

// FatalS - print error and exit program # panic (string only)
func FatalS(s string) {
	switch logLogger {
	case PHUSLU_PRETTY, PHUSLU_JSON:
		phuslu.Fatal().Msg(s)
	default:
		utils.PrintFatal(s)
	}
	PanicNow()
}

// CheckErrFatal: checks for error and exits the program #panic
func CheckErrFatal(err error) {
	if err != nil {
		_, src := utils.GetCaller()
		ErrorS("--Panic--")
		ErrorS("Error : " + err.Error() + " 💥")
		ErrorS("Reason: " + utils.Excuse())
		ErrorS("Source: " + src)
		ErrorS("--Panic--")
		FatalS("bailing out ... 😵 ")
	}
}

func PanicNow() {
	panic("😱  Houston - we have a problem here ...💥")
}

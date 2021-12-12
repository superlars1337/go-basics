package utils

// unicode emoji list :
// https://unicode.org/emoji/charts/full-emoji-list.html

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const symbol = "░"

// PrintTitle - prints nicely formated title
func PrintTitle(title string) {
	length := len(title) + 4
	color.Set(color.FgHiGreen)

	fmt.Println()
	printLine(symbol, length)
	fmt.Println(symbol + " " + title + " " + symbol)
	printLine(symbol, length)
	fmt.Println()

	color.Unset()
}

// PrintSubTitle - prints nicely formated subtitle
func PrintSubTitle(subtitle string) {
	fmt.Println()
	fmt.Println(subtitle)
	printLine(symbol, len(subtitle))
}

// PrintType - prints the type of a interface
func PrintType(name string, obj interface{}) {
	fmt.Printf("%s is type %T\n", name, obj)
}

//  ------------------------------------------------------------------------
//     Print Messages - used via LOG package    -
//  ------------------------------------------------------------------------

// PrintMessage - prints nicely formated message
func PrintMessage(message string) {
	fmt.Println("💬 " + message)
}

// PrintSuccess - prints nicely formated success info
func PrintSuccess(success string) {
	color.Green("🏁 YAY - " + success)
}

// ReadLine - prints nicely formated readline input
func ReadLine(msg string) (string, error) {
	fmt.Printf("%s ", msg)
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	return string(line), err
}

// PrintError - prints nicely formated error message
func PrintError(error string) {
	color.Red("💣 " + error)
}

// PrintFatal - prints nicely formated fatal message
func PrintFatal(error string) {
	color.Red("😵 " + error)
}

// PrintWarning - prints nicely formated warn message
func PrintWarning(warning string) {
	color.Yellow("🤭 " + warning)
}

// PrintInformation - prints nicely formated info message
func PrintInformation(info string) {
	color.Blue("❕ " + info)
}

// PrintDebug - prints nicely formated debug message
func PrintDebug(debug string) {
	color.Cyan("👉 " + debug)
}

//  ------------------------------------------------------------------------
//        Internal
//  ------------------------------------------------------------------------

func printLine(char string, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(char)
	}
	fmt.Println()
}

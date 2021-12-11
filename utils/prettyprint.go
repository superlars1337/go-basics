package utils

// unicode emoji list :
// https://unicode.org/emoji/charts/full-emoji-list.html

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// PrintTitle - prints nicely formated title
func PrintTitle(title string) {
	length := len(title) + 4
	symbol := "░"
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
	printLine("-", len(subtitle))
}

// PrintMessage - prints nicely formated message
func PrintMessage(message string) {
	color.White("💬 " + message)
}

// PrintSuccess - prints nicely formated success info
func PrintSuccess(success string) {
	color.Green("🏁 " + success)
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

// 	fmt.Printf("tm is type %T\n", tm)

//  ------------------------------------------------------------------------
//        Internal
//  ------------------------------------------------------------------------

func printLine(char string, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(char)
	}
	fmt.Println()
}

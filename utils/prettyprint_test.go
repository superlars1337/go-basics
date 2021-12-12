package utils

func ExamplePrintTitle() {
	PrintTitle("test")
	// Output:
	//░░░░░░░░
	//░ test ░
	//░░░░░░░░
}

func ExamplePrintSubTitle() {
	PrintSubTitle("test")
	// Output:
	//test
	//░░░░
}

func ExamplePrintMessage() {
	PrintMessage("test")
	// Output:
	//💬 test
}


package console

import (
	"fmt"
)

//PrintError : for printing error message
func PrintError(message interface{}) {
	fmt.Println("\033[31m", message, "\033[0m")
}

//PrintSuccess : for printing success message
func PrintSuccess(message interface{}) {
	fmt.Println("\033[1;88m", message, "\033[0m")
}

//PrintWarning : for printing Warning message
func PrintWarning(message interface{}) {
	fmt.Println("\033[1;91m", message, "\033[0m")
}

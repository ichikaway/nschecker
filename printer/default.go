package printer

import (
	"fmt"
)

var Printf = PrinterDefault
var ErrorPrintf = ErrorPrint

func PrinterDefault(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func PrinterNothing(format string, a ...interface{}) {
}

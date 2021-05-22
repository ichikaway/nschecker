package printer

import (
	"fmt"
)

var Printf = DefaultPrintf
var ErrorPrintf = ErrorPrint

func DefaultPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func PrinterNothing(format string, a ...interface{}) {
}

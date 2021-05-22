package printer

import (
	"fmt"
)

var handlePrintf = defaultPrintf

func Printf(format string, a ...interface{}) {
	handlePrintf(format, a...)
}

func SilentModeOn() {
	handlePrintf = printerNothing
}

func defaultPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func printerNothing(format string, a ...interface{}) {
}

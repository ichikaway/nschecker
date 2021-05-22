package printer

import (
	"fmt"
	"os"
)

func ErrorPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

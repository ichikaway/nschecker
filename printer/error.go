package printer

import (
	"fmt"
	"os"
)

func ErrorPrint(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

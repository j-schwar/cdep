package kernel

import (
	"fmt"
	"os"
)

func fatalError(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

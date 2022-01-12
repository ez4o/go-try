package try

import (
	"fmt"
	"runtime"
)

type StackInfo struct {
	pc   uintptr
	file string
	line int
}

func (si *StackInfo) String() string {
	f := runtime.FuncForPC(si.pc)
	return fmt.Sprintf("%s\n\t%s:%d\n", f.Name(), si.file, si.line)
}

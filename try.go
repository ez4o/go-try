package try

import (
	"fmt"
	"runtime"
)

func getStackTrace() string {
	stackTrace := ""

	// Skip caller [0, 1, 2, 3, 5, 6].
	for i := 4; ; i++ {
		if i == 5 || i == 6 {
			continue
		}

		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		f := runtime.FuncForPC(pc)
		stackTrace += fmt.Sprintf("%s\n\t%s:%d\n", f.Name(), file, line)
	}

	return stackTrace[:len(stackTrace)-1]
}

func Try(f func()) *Handler {
	var h *Handler

	func() {
		defer func() {
			if r := recover(); r != nil {
				h = &Handler{&Exception{r, getStackTrace()}}
			}
		}()

		f()
	}()

	return h
}

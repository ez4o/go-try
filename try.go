package try

import (
	"fmt"
	"runtime"
)

func getStackTrace(skip int, depth int) string {
	stackTrace := ""
	pc := make([]uintptr, depth)
	entriesAmount := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:entriesAmount])

	for {
		frame, more := frames.Next()
		stackTrace += fmt.Sprintf("%s\n\t%s:%d", frame.Func.Name(), frame.File, frame.Line)

		if more {
			stackTrace += "\n"
		} else {
			break
		}
	}

	return stackTrace
}

func Try(f func()) *Handler {
	const TRACE_DEPTH = 10

	var h *Handler

	stackTrace := getStackTrace(3, TRACE_DEPTH)

	func() {
		defer func() {
			if r := recover(); r != nil {
				h = &Handler{&Exception{r, getStackTrace(5, 1) + "\n" + stackTrace}}
			}
		}()

		f()
	}()

	return h
}

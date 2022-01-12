package try

import (
	"fmt"
	"runtime"
	"sync"
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

func Try(f func(func(any))) *Handler {
	const TRACE_DEPTH = 10

	var wg sync.WaitGroup
	var h *Handler

	stackTrace := getStackTrace(3, TRACE_DEPTH)

	wg.Add(1)
	go func() {
		defer wg.Done()

		f(func(err any) {
			if err != nil {
				h = &Handler{&Exception{err, getStackTrace(3, 1) + "\n" + stackTrace}}
				runtime.Goexit()
			}
		})
	}()

	wg.Wait()
	return h
}

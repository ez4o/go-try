package try

import (
	"runtime"
)

func getStackTrace() *StackTrace {
	stackInfos := make([]StackInfo, 0)

	// Skip caller [0, 1, 2, 3, 5, 6].
	for i := 4; ; i++ {
		if i == 5 || i == 6 {
			continue
		}

		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		stackInfos = append(stackInfos, StackInfo{pc, file, line})
	}

	return &StackTrace{stackInfos}
}

func Try(f func()) *Exception {
	var e *Exception

	func() {
		defer func() {
			if r := recover(); r != nil {
				e = &Exception{r, getStackTrace()}
			}
		}()

		f()
	}()

	return e
}

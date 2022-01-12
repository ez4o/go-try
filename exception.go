package try

import "fmt"

type Exception struct {
	Error      any
	stackTrace []StackInfo
}

func (e *Exception) PrintStackTrace() {
	stackTrace := ""

	for _, s := range e.stackTrace {
		stackTrace += s.String()
	}

	fmt.Println(stackTrace[:len(stackTrace)-1])
}

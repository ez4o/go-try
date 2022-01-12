package try

import "fmt"

type StackTrace struct {
	stackInfos []StackInfo
}

func (st *StackTrace) Print() {
	stackTrace := ""

	for _, si := range st.stackInfos {
		stackTrace += si.String()
	}

	fmt.Println(stackTrace[:len(stackTrace)-1])
}

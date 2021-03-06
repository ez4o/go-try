package main

import (
	"fmt"

	. "github.com/ez4o/go-try"
)

type Test struct {
	val int
}

func main() {
	Try(func() {
		ThrowOnError(1)
	}).Catch(func(s string) {
		fmt.Println("This is a string exception:", s)
	}).Catch(func(i int) {
		fmt.Println("This is an int exception:", i)
	})

	Try(func() {
		ThrowOnError("1")
	}).Catch(func(s string) {
		fmt.Println("This is a string exception:", s)
	}).Catch(func(i int) {
		fmt.Println("This is an int exception:", i)
	})

	Try(func() {
		ThrowOnError(Test{1})
	}).Catch(func(t Test) {
		fmt.Println("This is a Test exception:", t)
	})
}

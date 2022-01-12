package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/ez4o/go-try"
)

func main() {
	Try(func() {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		ThrowOnError(err)

		resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		ThrowOnError(err)

		var data []map[string]interface{}
		err = json.Unmarshal(b, &data)
		ThrowOnError(err)

		fmt.Println(data)
	}).Catch(func(e *Exception) {
		fmt.Println(e.Error)
		e.PrintStackTrace()
	})
}

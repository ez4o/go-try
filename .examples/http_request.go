package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/ez4o/go-try"
)

func main() {
	Try(func(throwOnError func(err any)) {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		throwOnError(err)

		resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		throwOnError(err)

		var data []map[string]interface{}
		err = json.Unmarshal(b, &data)
		throwOnError(err)

		fmt.Println(data)
	}).Catch(func(e *Exception) {
		fmt.Println(e.Error)
		fmt.Println(e.StackTrace)
	})
}

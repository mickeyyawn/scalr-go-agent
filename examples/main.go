package main

import (
	"fmt"

	scalyr "github.com/mickeyyawn/scalyr-go-agent"
)

func main() {

	fmt.Println("about to exercise the scalyr agent...")

	m := map[string]interface{}{"a": "apple", "b": 2}
	// or
	z := map[string]interface{}{}
	z["c"] = "this is c"
	z["d"] = "pear"
	// or
	h := map[string]interface{}{}
	h["f"] = "this is f"
	h["message"] = "my message is here"

	// or

	l := map[string]interface{}{"a": "banana", "b": 22, "message": "my most awesome message is here...."}

	scalyr.Event(Warning, m)
	Event(Info, z)
	Event(Info, h)
	Event(Info, l)

	for i := 0; i < 3; i++ {
		Event(Info, l)
	}

}

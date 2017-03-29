package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("about to exercise the scalyr agent...")
	//
	// for testing purposes let's grab the key from an env var
	//
	key := os.Getenv("SCALYR_KEY")

	//
	// let's mock up some events we want to push into scalyr
	//
	// here is a one line map with a few key/values
	oneLiner := map[string]interface{}{"a": "apple", "b": 2}

	//
	// here is an example with different syntax for setting the map values
	//
	// NOTE: if you don't set message, then these key/values will appear
	//       in the main line in the Scalyr UI
	//
	multipleLines := map[string]interface{}{}
	multipleLines["c"] = "this is c"
	multipleLines["d"] = "pear"

	//
	// now this is important. IF you specify message in the map, then that
	// line will show in the scalyr ui, and the other values will be "additional" attributes...
	//
	includeMessage := map[string]interface{}{}
	includeMessage["f"] = "this is f"
	includeMessage["message"] = "my message is here"

	//
	// and just another example of setting everything on one line...
	//

	oneLineWithMessage := map[string]interface{}{"a": "banana", "b": 22, "message": "my most awesome message is here...."}

	//
	// ok, let's pass those events into Scalyr, note we are
	// using various severity types... And we are looping
	// just so you get a few different events to look at
	// in the UI.

	for i := 0; i < 3; i++ {
		Event(Warning, oneLiner)
	}

	for i := 0; i < 3; i++ {
		Event(Info, multipleLines)
	}

	for i := 0; i < 3; i++ {
		Event(Info, includeMessage)
	}

	for i := 0; i < 3; i++ {
		Event(Error, oneLineWithMessage)
	}

	// Now log into your Scalyr account and you will see your events.

	fmt.Println("done exercising the scalyr agent...")

}

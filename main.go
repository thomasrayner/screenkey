package main

import (
	"fmt"
	"time"

	hook "github.com/robotn/gohook"
)

func main() {
	s := hook.Start()
	defer hook.End()

	done := false
	for !done {
		select {
		case i := <-s:
			if i.Kind == hook.KeyDown {
				fmt.Print(string(i.Keychar))
				time.Sleep(100 * time.Millisecond)
			}

		}
	}

}

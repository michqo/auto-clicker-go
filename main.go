package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	go addClicker()
	addHooks()
}

var clicking bool = false

func addHooks() {
	hook.Register(hook.KeyDown, []string{"'"}, func(e hook.Event) {
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"c"}, func(e hook.Event) {
		clicking = !clicking
	})

	s := hook.Start()
	<-hook.Process(s)
}

func addClicker() {
	const MIN int = 45
	const MAX int = 85

	for {
		if clicking {
			for {
				if !clicking {
					break
				}
				robotgo.Click()
				delay := rand.Intn(MAX-MIN) + MIN
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

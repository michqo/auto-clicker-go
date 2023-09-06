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

const MIN int = 40
const MAX int = 70
const IDLE int = 80

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

		time.Sleep(time.Duration(IDLE) * time.Millisecond)
	}
}

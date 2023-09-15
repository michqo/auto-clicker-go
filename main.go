package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var clicking bool

const MIN int = 35
const MAX int = 75

func main() {
	clickCh := make(chan bool)
	go watch(clickCh)
	addHooks(clickCh)
}

func addHooks(clickCh chan<- bool) {
	hook.Register(hook.KeyDown, []string{"'"}, func(e hook.Event) {
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"c"}, func(e hook.Event) {
		clickCh <- true
	})

	s := hook.Start()
	<-hook.Process(s)
}

func watch(clickCh <-chan bool) {
	for {
		<-clickCh
		clicking = !clicking
		if clicking {
			go clickMouse()
		}
	}
}

func clickMouse() {
	var delay int
	for {
		if !clicking {
			return
		}
		robotgo.Click()
		delay = rand.Intn(MAX-MIN) + MIN
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

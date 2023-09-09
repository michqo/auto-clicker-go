package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var clicking bool
var clickerRunning bool

const MIN int = 40
const MAX int = 70

func main() {
	clickCh := make(chan bool)
	go addClicker(clickCh)
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

func addClicker(clickCh <-chan bool) {
	for {
		<-clickCh
		clicking = !clicking
		if !clickerRunning && clicking {
			go clickMouse()
			clickerRunning = true
		}
	}
}

func clickMouse() {
	var delay int
	for {
		if !clicking {
			clickerRunning = false
			return
		}
		robotgo.Click()
		delay = rand.Intn(MAX-MIN) + MIN
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

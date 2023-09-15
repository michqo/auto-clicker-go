package main

import (
	hook "github.com/robotn/gohook"
)

var leftClicker Clicker = Clicker{delay: Delay{MIN: 35, MAX: 75}, running: false}

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
		leftClicker.running = !leftClicker.running
		if leftClicker.running {
			go leftClicker.activate()
		}
	}
}

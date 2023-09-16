package main

import (
	hook "github.com/robotn/gohook"
)

const (
	LEFT  = 0
	RIGHT = 1
)

var leftClicker Clicker = Clicker{delay: Delay{MIN: 35, MAX: 75}, running: false, button: "left"}
var rightClicker Clicker = Clicker{delay: Delay{MIN: 20, MAX: 45}, running: false, button: "right"}

func main() {
	clickCh := make(chan int)
	go watch(clickCh)
	addHooks(clickCh)
}

func addHooks(clickCh chan<- int) {
	hook.Register(hook.KeyDown, []string{"'"}, func(e hook.Event) {
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"c"}, func(e hook.Event) {
		clickCh <- LEFT
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		clickCh <- RIGHT
	})

	s := hook.Start()
	<-hook.Process(s)
}

func watch(clickCh <-chan int) {
	for {
		switch <-clickCh {
		case LEFT:
			leftClicker.running = !leftClicker.running
			if leftClicker.running {
				go leftClicker.activate()
			}
		case RIGHT:
			rightClicker.running = !rightClicker.running
			if rightClicker.running {
				go rightClicker.activate()
			}
		}

	}
}

package main

import (
	hook "github.com/robotn/gohook"
)

const (
	LEFT  = 0
	RIGHT = 1
)

var LEFT_DELAY Delay = Delay{MIN: 35, MAX: 75}
var RIGHT_DELAY Delay = Delay{MIN: 15, MAX: 30}

var leftClicker Clicker = Clicker{delay: LEFT_DELAY, running: false, button: "left"}
var rightClicker Clicker = Clicker{delay: RIGHT_DELAY, running: false, button: "right"}

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

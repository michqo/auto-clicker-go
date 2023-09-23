package main

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	hook "github.com/robotn/gohook"
)

const (
	LEFT  = 0
	RIGHT = 1
)

var leftClicker Clicker
var rightClicker Clicker

func main() {
	var cfg Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	leftClicker = Clicker{delay: cfg.LeftDelay, running: false, button: "left"}
	rightClicker = Clicker{delay: cfg.RightDelay, running: false, button: "right"}

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

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

var leftDeactivate []string
var rightDeactivate []string
var deactivate []string

func main() {
	loadConfig()
	clickCh := make(chan int)
	go watch(clickCh)
	addHooks(clickCh)
}

func loadConfig() {
	var cfg Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	leftClicker = Clicker{delay: cfg.LeftClick.Delay, running: false, button: "left"}
	rightClicker = Clicker{delay: cfg.RightClick.Delay, running: false, button: "right"}
	leftDeactivate = []string{cfg.LeftClick.Deactivate}
	rightDeactivate = []string{cfg.RightClick.Deactivate}
	deactivate = []string{cfg.Deactivate}
}

func addHooks(clickCh chan<- int) {
	hook.Register(hook.KeyDown, deactivate, func(e hook.Event) {
		hook.End()
	})

	hook.Register(hook.KeyDown, leftDeactivate, func(e hook.Event) {
		clickCh <- LEFT
	})

	hook.Register(hook.KeyDown, rightDeactivate, func(e hook.Event) {
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

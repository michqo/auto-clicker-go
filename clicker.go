package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type Delay struct {
	MIN int
	MAX int
}

type Clicker struct {
	delay   Delay
	running bool
}

func (c *Clicker) activate() {
	var delay int
	for {
		if !c.running {
			return
		}
		robotgo.Click()
		delay = rand.Intn(c.delay.MAX-c.delay.MIN) + c.delay.MIN
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

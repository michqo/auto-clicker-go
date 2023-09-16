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
	button  string
}

func (c *Clicker) activate() {
	var delay int
	var diff int = c.delay.MAX - c.delay.MIN
	for {
		if !c.running {
			return
		}
		robotgo.Click(c.button)
		delay = rand.Intn(diff) + c.delay.MIN
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

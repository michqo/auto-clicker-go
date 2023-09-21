package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type Delay struct {
	value     int
	threshold int
}

type Clicker struct {
	delay   Delay
	running bool
	button  string
}

func (c *Clicker) activate() {
	var delay int
	lower := c.delay.value - c.delay.threshold
	higher := c.delay.value + c.delay.threshold
	diff := higher - lower
	for {
		if !c.running {
			return
		}
		robotgo.Click(c.button)
		delay = rand.Intn(diff) + lower
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

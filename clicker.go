package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type Config struct {
	LeftDelay  Delay `yaml:"left"`
	RightDelay Delay `yaml:"right"`
}

type Delay struct {
	Value     int `yaml:"value"`
	Threshold int `yaml:"threshold"`
}

type Clicker struct {
	delay   Delay
	running bool
	button  string
}

func (c *Clicker) activate() {
	var delay int
	lower := c.delay.Value - c.delay.Threshold
	higher := c.delay.Value + c.delay.Threshold
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

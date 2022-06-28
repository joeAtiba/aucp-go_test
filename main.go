package main

import (
	"fmt"
	"os"
	"strings"
)

type device interface {
	turnOn() string
	update(version float32)
}

type iPhone struct {
	name    string
	model   string
	version float32
}

type iMac struct {
	name    string
	model   string
	version float32
}

func (phone iPhone) turnOn() string {
	return "iOS starting up..."
}

func (mac iMac) turnOn() string {
	return "macOS starting up..."
}

func (phone *iPhone) update(v float32) {
	phone.version = v
}

func (mac *iMac) update(v float32) {
	mac.version = v
}

func main() {
	rc := 0
	defer func() { os.Exit(rc) }()

	dev1 := iPhone{"iPhone", "17 Pro", 13.1}
	dev2 := iMac{"iMac", "27 15k Retina", 10.15}

	fmt.Println(dev1.turnOn())
	fmt.Println(dev2.turnOn())

	devices := []device{&dev1, &dev2}

	for _, dev := range devices {
		if strings.Contains(dev.turnOn(), "iOS") {
			dev.update(14.0)
		} else if strings.Contains(dev.turnOn(), "macOS") {
			dev.update(11.00)
		}
	}

	fmt.Println(dev1)
	fmt.Println(dev2)
}

package main

import "time"

type Thing struct {
	Name         string
	Associations []Thing
}

type Agent struct {
	Thing
	Intention Action // temporary short term goal
	Memories  map[time.Time]Event
	Mood      int
}

type Action struct {
	Description string
	Instigator  *Agent
	Expectation string
	Receiver    *Agent
}

func (*Agent) Action() {
}

type Location struct {
	Thing
	Geo string // marker
}

type Event struct {
	What  Action
	Who   []Agent
	Where Location
}

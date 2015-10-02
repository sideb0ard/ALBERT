package main

type Thing struct {
	Name         string
	Associations []Thing
}

type Event struct {
	What string
}

package main

type Asteroid struct {
	Location Point
	CanSee   []*Asteroid
}

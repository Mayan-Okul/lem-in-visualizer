package main

type Room struct {
	Name    string
	X, Y    int
	IsStart bool
	IsEnd   bool
}

type Link struct {
	From, To string
}

type Move struct {
	Ant  string
	Room string
}

type Turn struct {
	Moves []Move
}

type Colony struct {
	NumAnts int
	Rooms   []Room
	Links   []Link
	Turns   []Turn
}
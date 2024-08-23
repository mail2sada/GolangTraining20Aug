package main

import "fmt"

type Player interface {
	GetStats() string
	GetBest() int
}

type Cricker struct {
	Role    string
	Best    int
	Matches int
	Avg     int
}

func (c Cricker) GetStats() string {
	return fmt.Sprintln("Cricker:", c.Matches, ":", c.Best, ":", c.Avg)
}

func (c Cricker) GetBest() int {
	return c.Best
}

type HockeyPlayer struct {
	Role    string
	Matches int
	Goals   int
}

func (c HockeyPlayer) GetStats() string {
	return fmt.Sprintln("HockeyPlayer:", c.Matches, ":", c.Goals, ":")
}

func (c HockeyPlayer) GetBest() int {
	return c.Goals
}

func PrintPlayerStats(p Player) {
	fmt.Println(p.GetStats())
	fmt.Println(p.GetBest())
}

func main() {
	fmt.Println("Interfaces")

	var p Player = Cricker{}
	var p1 Player = HockeyPlayer{}

	PrintPlayerStats(p)
	PrintPlayerStats(p1)
}

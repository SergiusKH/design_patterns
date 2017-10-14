package main

import (
	"fmt"
	"time"
)

func main() {
	var maker WatchMaker = new(RomeWatchMaker)
	var watch Watch = maker.CreateWatch()
	watch.ShowTime()
}

type Watch interface {
	ShowTime()
}

type DigitalWatch struct{}

func (d *DigitalWatch) ShowTime() {
	fmt.Println(time.Now())
}

type RomeWatch struct{}

func (r *RomeWatch) ShowTime() {
	fmt.Println("VII-XX")
}

type WatchMaker interface {
	CreateWatch() Watch
}

type DigitalWatchMaker struct{}

func (d *DigitalWatchMaker) CreateWatch() Watch {
	return new(DigitalWatch)
}

type RomeWatchMaker struct{}

func (r *RomeWatchMaker) CreateWatch() Watch {
	return new(RomeWatch)
}

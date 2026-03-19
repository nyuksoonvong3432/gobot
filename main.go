package main

import (
	"log"
	"time"
)

type orderType int

const (
	Normal orderType = iota
	Vip
)

type orderStatus int

const (
	Idle orderStatus = iota
	Processing
	Completed
)

type order struct {
	id        int
	orderType orderType
	status    orderStatus
}

type app struct {
	botId               int
	botCount            int
	processSecondPerBot time.Duration
}

func main() {
	app := app{
		botId:               1,
		botCount:            5,
		processSecondPerBot: time.Second * 10,
	}

	results := make(chan order)
	orders := make(chan order)
	app.subscribe(orders, results)

	orders <- order{id: 1, orderType: Normal, status: Idle}
	orders <- order{id: 2, orderType: Normal, status: Idle}
	orders <- order{id: 3, orderType: Normal, status: Idle}
	close(orders)

	for r := range results {
		log.Printf("Processed order %v", r)
	}
}

func (a *app) subscribe(orders chan order, results chan order) {
	go func() {
		for o := range orders {
			go a.assignOrderToBot(&o, results)
			a.botId++
		}
	}()
}

func (a *app) assignOrderToBot(o *order, results chan<- order) {
	log.Printf("Processing order %d with bot %v ...", o.id, a.botId)
	time.Sleep(a.processSecondPerBot)
	log.Printf("Finished processing order %d with bot %v", o.id, a.botId)
	o.status = Completed
	results <- *o
}

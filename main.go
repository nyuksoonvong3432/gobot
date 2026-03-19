package main

import (
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
	botCount            int
	processSecondPerBot time.Duration
}

type bot struct {
	id                  int
	order               *order
	processSecondPerBot time.Duration
}

func main() {
	app := app{
		botCount:            5,
		processSecondPerBot: time.Second * 10,
	}

	bots := make(chan bot)
	results := make(chan order)
	orders := make(chan order)

	app.subscribe(orders, results, bots)

	for i := 1; i <= 10; i++ {
		orders <- order{id: i, orderType: Normal, status: Idle}
	}

	for i := 1; i <= 3; i++ {
		bots <- bot{id: i, processSecondPerBot: app.processSecondPerBot}
	}

	for i := 10; i <= 13; i++ {
		orders <- order{id: i, orderType: Vip, status: Idle}
	}

	close(orders)
	// for r := range results {
	// 	log.Printf("Processed order %v", r)
	// }
}

func (a *app) subscribe(orders chan order, results chan order, bots chan bot) {
	go func() {
		for order := range orders {
			b := bot{
				order:               &order,
				processSecondPerBot: a.processSecondPerBot,
			}
			go b.processOrder(&order, results)
		}
	}()
}

func (b *bot) processOrder(o *order, results chan<- order) {
	o.status = Processing
	time.Sleep(b.processSecondPerBot)
	o.status = Completed
	results <- *o
	b.order = nil
}

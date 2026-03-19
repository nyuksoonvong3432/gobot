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

type order struct {
	id        int
	orderType orderType
}

type app struct {
	botCount            int
	processSecondPerBot time.Duration
}

func main() {
	app := app{
		botCount:            5,
		processSecondPerBot: time.Second * 10,
	}

	orders := make(chan order)
	app.subscribe(orders)

	orders <- order{id: 1, orderType: Normal}
	orders <- order{id: 2, orderType: Normal}
	orders <- order{id: 3, orderType: Normal}
	close(orders)
}

func (a *app) subscribe(orders <-chan order) {
	go func() {
		for o := range orders {
			log.Printf("Received order id: %d", o.id)

		}
	}()
}

func bot(id int, orders chan<- order) {
	log.Println("Bot ", id)
	time.Sleep(10)
}

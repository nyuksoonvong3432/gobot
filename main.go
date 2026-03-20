package main

import (
	"fmt"
	"sync"
	"time"
)

type orderType int

const (
	Normal orderType = iota
	Vip
)

func (ot orderType) String() string {
	switch ot {
	case Normal:
		return "Normal"
	case Vip:
		return "Vip"
	default:
		return "Unknown"
	}
}

type orderStatus int

const (
	Idle orderStatus = iota
	Processing
	Completed
)

func (os orderStatus) String() string {
	switch os {
	case Idle:
		return "Idle"
	case Processing:
		return "Processing"
	case Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}

type order struct {
	id        int
	orderType orderType
	status    orderStatus
}

func (o order) String() string {
	return fmt.Sprintf("Order{id: %d, orderType: %s, status: %s}", o.id, o.orderType, o.status)
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

	var wg sync.WaitGroup
	app.start(orders, &wg)

	orders <- order{id: 1, orderType: Normal, status: Idle}
	orders <- order{id: 2, orderType: Normal, status: Idle}
	orders <- order{id: 3, orderType: Normal, status: Idle}
	close(orders)

	wg.Wait()

	fmt.Println("App exited.")
}

func (a *app) start(orders <-chan order, wg *sync.WaitGroup) {
	for b := 1; b <= a.botCount; b++ {
		wg.Go(func() {
			bot(b, orders, a.processSecondPerBot)
		})
	}
}

func bot(id int, orders <-chan order, processSecond time.Duration) {
	for order := range orders {
		fmt.Printf("Bot %d start processing order %v\n", id, order)
		order.status = Processing
		time.Sleep(processSecond)
		order.status = Completed
		fmt.Printf("Bot %d completed order %v\n", id, order)
	}
}

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

	normalOrders := make(chan order)
	vipOrders := make(chan order)

	var wg sync.WaitGroup
	app.start(normalOrders, vipOrders, &wg)

	normalOrders <- order{id: 1, orderType: Normal, status: Idle}
	normalOrders <- order{id: 2, orderType: Normal, status: Idle}
	normalOrders <- order{id: 3, orderType: Normal, status: Idle}
	normalOrders <- order{id: 4, orderType: Normal, status: Idle}
	normalOrders <- order{id: 5, orderType: Normal, status: Idle}

	normalOrders <- order{id: 6, orderType: Normal, status: Idle}
	vipOrders <- order{id: 7, orderType: Vip, status: Idle}
	vipOrders <- order{id: 8, orderType: Vip, status: Idle}

	close(vipOrders)
	close(normalOrders)

	wg.Wait()

	fmt.Println("App exited.")
}

func (a *app) start(normalOrders <-chan order, vipOrders <-chan order, wg *sync.WaitGroup) {
	for b := 1; b <= a.botCount; b++ {
		wg.Go(func() {
			bot(b, normalOrders, vipOrders, a.processSecondPerBot)
		})
	}
}

func bot(id int, normalOrders <-chan order, vipOrders <-chan order, processSecond time.Duration) {
	for {
		select {
		case o, more := <-vipOrders:
			if !more {
				return
			}
			processOrder(id, &o, processSecond)
		case o, more := <-normalOrders:
			if !more {
				return
			}
			processOrder(id, &o, processSecond)
		}
	}
}

func processOrder(botId int, o *order, processSecond time.Duration) {
	fmt.Printf("Bot %d processing order %v\n", botId, o)
	o.status = Processing
	time.Sleep(processSecond)
	fmt.Printf("Bot %d completed order %v\n", botId, o)
	o.status = Completed
}

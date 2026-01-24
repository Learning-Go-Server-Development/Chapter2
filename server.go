package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Customer struct {
	Name           string
	Address        string
	CustomerNumber string
	Phone          string
}

type StoreOrder interface {
	CreateOrder(cid string, sku string)
	UpdateOrder(sku string)
	Print(label string)
}

type Order struct {
	CustomerNumber string
	ItemSku        string
	TimeCreated    time.Time
	UpdatedTime    time.Time
}

func (o *Order) CreateOrder(cid string, sku string) {
	o.CustomerNumber = cid
	o.ItemSku = sku
	o.TimeCreated = time.Now()
}

func (o *Order) UpdateOrder(sku string) {
	o.ItemSku = sku
	o.UpdatedTime = time.Now()
}

func (o *Order) Print(label string) {
	fmt.Println(label, o)
}

func (o *Order) New() StoreOrder {
	return o
}

type DeliverOrder struct {
	CustomerNumber string
	ItemSku        string
	TimeCreated    time.Time
	UpdatedTime    time.Time
	DeliveryTime   time.Time
}

func (o *DeliverOrder) CreateOrder(cid string, sku string) {
	o.CustomerNumber = cid
	o.ItemSku = sku
	o.TimeCreated = time.Now()
}

func (o *DeliverOrder) UpdateOrder(sku string) {
	o.ItemSku = sku
	o.UpdatedTime = time.Now()
	o.DeliveryTime = time.Now().Add(time.Hour * 24 * 14)
}

func (o *DeliverOrder) Print(label string) {
	fmt.Println("This is a deliver order and will be delivered in two weeks or less "+label, o)
	if !o.DeliveryTime.IsZero() {
		fmt.Println("Delivery Date: ", o.DeliveryTime)
	}
}

func (o *DeliverOrder) New() StoreOrder {
	return o
}

func main() {

	// moved to an interface below
	// var odr Order
	// odr.CreateOrder("12345", "sku123")
	// fmt.Println("new order: ", odr)

	var odr Order
	so := odr.New()
	so.CreateOrder("12345", "sku123")
	so.Print("new order: ")

	so.UpdateOrder("sku123---part 2")
	so.Print("updated order: ")

	var dodr DeliverOrder
	dso := dodr.New()
	dso.CreateOrder("12345", "sku123888")
	dso.Print("(new order): ")

	dso.UpdateOrder("sku123888---part 2")
	dso.Print("(updated order): ")

	var wg sync.WaitGroup

	var msg1 string = "goroutine 1 = "
	msg2 := "goroutine 2 = "

	var ctl1 int = 10
	ctl2 := 20

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= ctl1; i++ {
			fmt.Println(msg1, i)
		}
	}()

	wg.Add(1)
	go func() {
		duration := time.Second
		time.Sleep(duration * 3)
		defer wg.Done()

		for i := 11; i <= ctl2; i++ {
			fmt.Println(msg2, i)
		}
	}()

	wg.Wait()

	port := "3000"
	msg := "Server starting on port "
	fmt.Println(msg + port)
	http.ListenAndServe(":3000", nil)
}

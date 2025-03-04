package adapters

import (
	"fmt"
	"hexagonal-architecture/core"
	"hexagonal-architecture/ports"
)

type CashOrderAdapter struct {
	LemonadeStand *core.LemonadeStand
}

func NewCashOrderAdapter(lemonadeStand *core.LemonadeStand) ports.OrderPort {
	return &CashOrderAdapter{LemonadeStand: lemonadeStand}
}

func (c *CashOrderAdapter) PlaceOrder(amount float64) {
	fmt.Println("🤑 Customer pays with Cash")
	fmt.Println(c.LemonadeStand.TakePayment(amount))
	fmt.Println(c.LemonadeStand.MakeLemonade())
}

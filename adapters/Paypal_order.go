package adapters

import (
	"fmt"
	"hexagonal-architecture/core"
	"hexagonal-architecture/ports"
)

type PayPalOrderAdapter struct {
	LemonadeStand *core.LemonadeStand
}

func NewPayPalOrderAdapter(lemonadeStand *core.LemonadeStand) ports.OrderPort {
	return &PayPalOrderAdapter{LemonadeStand: lemonadeStand}
}

func (p *PayPalOrderAdapter) PlaceOrder(amount float64) {
	fmt.Println("💻 Customer pays with PayPal")
	fmt.Println(p.LemonadeStand.TakePayment(amount))
	fmt.Println(p.LemonadeStand.MakeLemonade())
}

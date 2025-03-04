package adapters

import (
	"fmt"
	"hexagonal-architecture/core"
	"hexagonal-architecture/ports"
)

type GooglePayOrderAdapter struct {
	LemonadeStand *core.LemonadeStand
}

func NewGooglePayOrderAdapter(lemonadeStand *core.LemonadeStand) ports.OrderPort {
	return &GooglePayOrderAdapter{LemonadeStand: lemonadeStand}
}

func (g *GooglePayOrderAdapter) PlaceOrder(amount float64) {
	fmt.Println("📲 Customer pays using Google Pay")
	fmt.Println(g.LemonadeStand.TakePayment(amount))
	fmt.Println(g.LemonadeStand.MakeLemonade())
}

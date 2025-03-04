package adapters

import (
	"fmt"
	"hexagonal-architecture/core"
	"hexagonal-architecture/ports"
)

type MobileMoneyOrderAdapter struct {
	LemonadeStand *core.LemonadeStand
}

func NewMobileMoneyOrderAdapter(lemonadeStand *core.LemonadeStand) ports.OrderPort {
	return &MobileMoneyOrderAdapter{LemonadeStand: lemonadeStand}
}

func (m *MobileMoneyOrderAdapter) PlaceOrder(amount float64) {
	fmt.Println("📱 Customer pays using Mobile Money (M-Pesa)")
	fmt.Println(m.LemonadeStand.TakePayment(amount))
	fmt.Println(m.LemonadeStand.MakeLemonade())
}

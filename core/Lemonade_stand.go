package core

import "fmt"

// LemonadeStand represents our business
type LemonadeStand struct {
	Name string
}

func NewLemonadeStand(name string) *LemonadeStand {
	return &LemonadeStand{Name: name}
}

func (l *LemonadeStand) TakePayment(amount float64) string {
	return fmt.Sprintf("💰 Payment of $%.2f received at %s!", amount, l.Name)
}

func (l *LemonadeStand) MakeLemonade() string {
	return "🍋 Fresh Lemonade is ready!"
}

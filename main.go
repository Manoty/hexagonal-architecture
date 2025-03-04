package main

import (
	"fmt"
	"hexagonal-architecture/adapters"
	"hexagonal-architecture/core"
)

func main() {
	fmt.Println("🚀 Welcome to Lemonade Stand!")

	// Create Lemonade Stand Business
	lemonadeStand := core.NewLemonadeStand("Kevin's Lemonade Stand")

	// Cash Payment
	cashOrder := adapters.NewCashOrderAdapter(lemonadeStand)
	fmt.Println("\nProcessing Cash Order:")
	cashOrder.PlaceOrder(5.0)

	// PayPal Payment
	paypalOrder := adapters.NewPayPalOrderAdapter(lemonadeStand)
	fmt.Println("\nProcessing PayPal Order:")
	paypalOrder.PlaceOrder(10.0)

	// Google Pay Payment
	googlePayOrder := adapters.NewGooglePayOrderAdapter(lemonadeStand)
	fmt.Println("\nProcessing Google Pay Order:")
	googlePayOrder.PlaceOrder(15.0)

	// Mobile Money Payment (Newly Added)
	mobileMoneyOrder := adapters.NewMobileMoneyOrderAdapter(lemonadeStand)
	fmt.Println("\nProcessing Mobile Money Order:")
	mobileMoneyOrder.PlaceOrder(20.0)
}

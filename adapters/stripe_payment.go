package adapters

import (
	"fmt"
	"hexagonal-architecture/ports"
)

// StripePaymentAdapter implements the PaymentPort interface
type StripePaymentAdapter struct{}

// NewStripePaymentAdapter creates a new instance of StripePaymentAdapter
func NewStripePaymentAdapter() ports.PaymentPort {
	return &StripePaymentAdapter{}
}

// ProcessPayment simulates Stripe payment processing
func (s *StripePaymentAdapter) ProcessPayment(amount float64, currency string) bool {
	fmt.Printf("✅ [MOCK] Stripe: Processing payment of %.2f %s...\n", amount, currency)
	fmt.Println("✅ [MOCK] Payment successful!")
	return true
}

package ports

// PaymentPort defines the behavior of a payment adapter
type PaymentPort interface {
	ProcessPayment(amount float64, currency string) bool
}

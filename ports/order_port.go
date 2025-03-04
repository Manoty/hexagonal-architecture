package ports

type OrderPort interface {
	PlaceOrder(amount float64)
}
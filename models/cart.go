package models

const (
	// Status
	New  = "New"
	Paid = "Paid"
)

type Cart struct {
	CartId     string
	ProductId  string
	CustomerId string
	Status     string
}

package domain

type Deliverer interface {
	Deliver(pharmacy Pharmacy) error
}
package decorators

type Beverage interface {
	GetDescription() string
	Cost() float64
	GetSize() string
}

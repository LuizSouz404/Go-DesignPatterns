package decorators

type Mocha struct {
	Beverage Beverage
}

func (m *Mocha) GetDescription() string {
	return (m.Beverage.GetDescription() + ", Mocha")
}

func (m *Mocha) Cost() float64 {
	switch m.Beverage.GetSize() {
	case "small":
		return m.Beverage.Cost() + 0.5

	case "medium":
		return m.Beverage.Cost() + 0.7

	case "big":
		return m.Beverage.Cost() + 1.0

	default:
		return m.Beverage.Cost() + 0.7
	}
}

func (m *Mocha) GetSize() string {
	return m.Beverage.GetSize()
}

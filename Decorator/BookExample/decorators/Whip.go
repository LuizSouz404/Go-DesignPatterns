package decorators

type Whip struct {
	Beverage Beverage
}

func (m *Whip) GetDescription() string {
	return (m.Beverage.GetDescription() + ", Whip")
}

func (m *Whip) Cost() float64 {
	switch m.Beverage.GetSize() {
	case "small":
		return m.Beverage.Cost() + 0.25

	case "medium":
		return m.Beverage.Cost() + 0.45

	case "big":
		return m.Beverage.Cost() + 0.75

	default:
		return m.Beverage.Cost() + 0.45
	}
}

func (w *Whip) GetSize() string {
	return w.Beverage.GetSize()
}

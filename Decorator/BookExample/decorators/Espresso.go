package decorators

type Espresso struct {
	Size string
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) Cost() float64 {
	return 1.5
}

func (e *Espresso) GetSize() string {
	return e.Size
}
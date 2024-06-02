package units

type System int

const (
	Metric System = iota
	Freedom
)

type Unit struct {
	Name   string
	System System
}

type Converter struct {
	Units []Unit
}

func NewConverter() Converter {
	return Converter{
		Units: []Unit{
			{Name: "g", System: Metric},
			{Name: "kg", System: Metric},
			{Name: "ml", System: Metric},
			{Name: "l", System: Metric},
			{Name: "oz", System: Freedom},
			{Name: "lb", System: Freedom},
			{Name: "fl oz", System: Freedom},
			{Name: "pt", System: Freedom},
		},
	}
}

func (converter Converter) Convert(amount float64, from, to string) (float64, error) {
	return 0, nil
}

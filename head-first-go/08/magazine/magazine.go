// B"H

package magazine

type Subscriber struct {
	Name      string
	Rate      float64
	Active    bool
	Magazines []Magazine
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	id         int
	Street     string
	City       string
	State      string
	PostalCode string
}

type Magazine struct {
	MagazineTitle string
	Price         float64
}

func (m *Magazine) ApplyDiscount(discountPct float64) {
	m.Price = m.Price - (m.Price * discountPct)
}

func ApplyMagDiscount(m Magazine, discountPct float64) Magazine {
	m.Price = m.Price - (m.Price * discountPct)

	return m
}

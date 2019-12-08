// B"H

package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
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

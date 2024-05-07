package misc

type Sort interface {
	GetName() string
	IsAscending() bool
	GetOrder() int
}

type sort struct {
	name  string
	asc   bool
	order int
}

func NewSort(name string, asc bool, order int) Sort {
	return sort{name: name, asc: asc, order: order}
}

func (s sort) GetName() string {
	return s.name
}

func (s sort) IsAscending() bool {
	return s.asc
}

func (s sort) GetOrder() int {
	return s.order
}

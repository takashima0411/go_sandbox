package money

type Gosen struct {
	amount int64
}

func NewGosen() Gosen {
	return Gosen{5000}
}

func (g Gosen) Price() int64 {
	return g.amount
}

func (g Gosen) Add(m Money) Money {
	return NewVariable(g.amount + m.Price())
}

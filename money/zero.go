package money

type Zero struct {
	amount int64
}

func NewZero() Zero {
	return Zero{0}
}

func (s Zero) Price() int64 {
	return s.amount
}

func (s Zero) Add(m Money) Money {
	return m
}

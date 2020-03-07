package money

type Sen struct {
	amount int64
}

func NewSen() Sen {
	return Sen{1000}
}

func (s Sen) Price() int64 {
	return 1000
}

func (s Sen) Add(m Money) Money {
	return NewVariable(s.amount + m.Price())
}

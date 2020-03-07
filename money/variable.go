package money

type Variable struct {
	amount int64
}

func NewVariable(v int64) Variable {
	return Variable{v}
}

func (v Variable) Price() int64 {
	return v.amount
}

func (v Variable) Add(m Money) Money {
	return NewVariable(v.amount + m.Price())
}

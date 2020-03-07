package money

type Money interface {
	Price() int64
	Add(m Money) Money
}

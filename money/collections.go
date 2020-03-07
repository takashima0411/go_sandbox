package money

func foldl(ls []Money, z Money, op func(Money, Money) Money) Money {
	if len(ls) == 0 {
		return z
	} else {
		return foldl(ls[1:], op(z, ls[0]), op)
	}
}


func Fold(ls []Money, z []Money, op func(Money, []Money) []Money) []Money {
	if len(ls) == 0 {
		return z
	} else {
		return Fold(ls[1:], op(ls[0], z), op)
	}
}

func Map(ls []Money, op func(Money) Money) []Money {
	return Fold(ls, nil, func(v Money, l []Money) []Money {
		return append(l, op(v))
	})
}

func Reverse(ls []Money) []Money {
	return Fold(ls, nil, func(v Money, l []Money) []Money {
		return append([]Money{v}, l...)
	})
}

package models

type A struct {
	id   int64
	name string
}

func NewA(id int64, name string) A {
	return A{id, name}
}

func (a A) GetID() int64 {
	return a.id
}

func (a A) GetName() string {
	return a.name
}

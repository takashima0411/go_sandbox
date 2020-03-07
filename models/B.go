package models

type B struct {
	id   int64
	name string
}

func NewB(id int64, name string) B {
	return B{id, name}
}

func (b B) GetID() int64 {
	return b.id
}

func (b B) GetName() string {
	return b.name
}

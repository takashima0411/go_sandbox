package models

type IF interface {
	GetID() int64
	GetName() string
}

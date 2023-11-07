package entity

import "github.com/google/uuid"

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Id:   uuid.New().String(),
		Name: name,
		Age:  age,
	}
}

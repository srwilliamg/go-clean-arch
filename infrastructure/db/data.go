package database

import (
	"go-clean-arch/entities"
)

type DataObject struct{}

func New() *DataObject {
	return &DataObject{}
}

var data = []*entities.User{
	{ID: 1, Name: "Pepe", Age: 3},
	{ID: 2, Name: "Juan", Age: 4},
	{ID: 3, Name: "Ramiro", Age: 5},
	{ID: 4, Name: "Pepe", Age: 6},
}

func (do *DataObject) MultipleUserByName(name string) []*entities.User {
	return data
}

// func Use() {
// 	do := New()
// 	ids := do.MultipleUserByName("Pepe")
// 	fmt.Printf("ids: %v\n", ids)
// }

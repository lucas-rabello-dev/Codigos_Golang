package models

import "fmt"

type Item struct {
	ID int
	Name string
	Quantity int
	Price float64
}


func (i Item) Info() string {
	// Sprintf -> para formatar strings sem imprimir elas na tela
	return fmt.Sprintf("ID : %d | Name: %s | Quantity: %d | Price: %.2f", i.ID, i.Name, i.Quantity, i.Price)
}

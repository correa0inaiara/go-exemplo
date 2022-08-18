package main

import "fmt"

type Car struct {
	Name string
	Model string
}

func (c Car) Andar() {
	fmt.Println("O Carro ", c.Name, " está andando.")
}

func main() {
	carro := Car{"Tucson", "Hyundai"}
	fmt.Println("\n", carro, "\nName: ", carro.Name, "\nModel: ", carro.Model)
	carro.Andar()

	carro.Name = "A3 Sedan"
	carro.Model = "Audi"
	fmt.Println("\n", carro, "\nName: ", carro.Name, "\nModel: ", carro.Model)
	carro.Andar()
}
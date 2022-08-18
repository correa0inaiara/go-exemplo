package main

import (
	"fmt"
)

func main() {
	// println("Hello, World!")
	// http.ListenAndServe(":8080", nil)

	var a int
	a = 10
	fmt.Println(a)

	nome := "Naiara" // criação e atribuição
	nome = "Naiara C" // string
	fmt.Println(nome)

	result, err := soma(7,2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// fmt.Println("\nsoma func\nresult: ", result, "\nerr: ", err)
}

func soma(a, b int) (int, error) {
	var result int
	result = a+b
	if result > 10 {
		return result, fmt.Errorf("soma é maior que 10")
	}
	return result, nil
}

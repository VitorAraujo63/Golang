package main

import {
	"fmt"
	"github.com/VitorAraujo63/Golang"
	"github.com/google/uuid"
}

func main() {
	S := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	
	fmt.Println(carro.andar())
	fmt.Println("Resultado: ", S)
	fmt.Println(matematica.A)
}
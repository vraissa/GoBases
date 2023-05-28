package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"
)

type AnimalFunc func(int) (float64, error)

func Animal(animalType string) (AnimalFunc, error) {
	switch animalType {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		return nil, errors.New("animal não encontrado")
	}
}

func dogFunc(quantity int) (float64, error) {
	foodAmount := 10.0 * float64(quantity)
	return foodAmount, nil
}

func catFunc(quantity int) (float64, error) {
	foodAmount := 5.0 * float64(quantity)
	return foodAmount, nil
}

func hamsterFunc(quantity int) (float64, error) {
	foodAmount := 0.25 * float64(quantity)
	return foodAmount, nil
}

func tarantulaFunc(quantity int) (float64, error) {
	foodAmount := 0.15 * float64(quantity)
	return foodAmount, nil
}

func main() {
	animalDog, _ := Animal(dog)
	animalCat, _ := Animal(cat)

	var amount float64
	amountDog, _ := animalDog(5)
	amountCat, _ := animalCat(8)
	amount += amountDog
	amount += amountCat

	fmt.Println("Quantidade total de comida:", amount, "kg")
}

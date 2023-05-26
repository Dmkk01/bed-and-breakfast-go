package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100, 0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is", result)
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

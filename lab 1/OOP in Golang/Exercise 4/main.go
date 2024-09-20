package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func ConvertToJSON(product Product) (string, error) {
	jsonData, err := json.Marshal(product)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DecodeFromJSON(jsonString string) (Product, error) {
	var product Product
	err := json.Unmarshal([]byte(jsonString), &product)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func main() {
	product := Product{
		Name:     "pen",
		Price:    999.99,
		Quantity: 10,
	}

	jsonStr, err := ConvertToJSON(product)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	} else {
		fmt.Println("JSON:", jsonStr)
	}

	decodedProduct, err := DecodeFromJSON(jsonStr)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	} else {
		fmt.Printf("Decoded Product: %+v\n", decodedProduct)
	}
}

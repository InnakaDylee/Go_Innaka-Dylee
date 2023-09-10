package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Buatlah sebuah program yang mengambil data dari sebuah REST API dengan ketentuan sebagai berikut:
// Menerapkan penggunaan concurrency (goroutine, channel).
// Endpoint yang digunakan: https://fakestoreapi.com/products
// Contoh output dari program yang dibuat adalah sebagai berikut:

var baseURL = "https://fakestoreapi.com/products"

type ProductData struct {
	Title    string		`json:"title"`
	Price    float64	`json:"price"`
	Category string		`json:"category"`
}

func fetchData() ([]ProductData, error) {
	var err error
	var client = http.Client{}
	var data []ProductData

	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var data, err = fetchData()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	fmt.Println("Products data")
	for _, each := range data {
		fmt.Println("===")
		fmt.Printf(" title 	  : %s \n price 	  : %f \n categoty : %s \n", each.Title, each.Price, each.Category)
		fmt.Println("===")
	}
}

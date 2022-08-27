package product

import "fmt"

type Product struct {
	ID    int
	Title string
}

var allProducts []Product = []Product{}

func init() {
	for i := 0; i < 30; i++ {
		allProducts = append(allProducts, Product{ID: i, Title: fmt.Sprintf("product_%d", i)})
	}
}

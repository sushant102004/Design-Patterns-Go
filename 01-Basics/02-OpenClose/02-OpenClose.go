/*
	This states that and module, type or a function is open for extension but closed for modifications.
*/

package main

import "fmt"

// Product is a struct representing a product with a name and price.
type Product struct {
	Name  string
	Price float64
}

// Filter is an interface for filtering products.
type Filter interface {
	Filter(products []Product) []Product
}

// PriceFilter is an implementation of the Filter interface that filters products by price.
type PriceFilter struct {
	MaxPrice float64
}

func (pf PriceFilter) Filter(products []Product) []Product {
	var filteredProducts []Product
	for _, product := range products {
		if product.Price <= pf.MaxPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

// NameFilter is an implementation of the Filter interface that filters products by name.
type NameFilter struct {
	SearchName string
}

func (nf NameFilter) Filter(products []Product) []Product {
	var filteredProducts []Product
	for _, product := range products {
		if product.Name == nf.SearchName {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

// ApplyFilter applies a filter to a list of products and returns the filtered list.
func ApplyFilter(filter Filter, products []Product) []Product {
	return filter.Filter(products)
}

func main() {
	products := []Product{
		{"Laptop", 1000.0},
		{"Smartphone", 500.0},
		{"Tablet", 300.0},
		{"Headphones", 100.0},
	}

	fmt.Println("All Products:")
	fmt.Println(products)

	priceFilter := PriceFilter{MaxPrice: 400.0}
	priceFilteredProducts := ApplyFilter(priceFilter, products)
	fmt.Println("\nProducts under $400:")
	fmt.Println(priceFilteredProducts)

	nameFilter := NameFilter{SearchName: "Laptop"}
	nameFilteredProducts := ApplyFilter(nameFilter, products)
	fmt.Println("\nProducts with the name 'Laptop':")
	fmt.Println(nameFilteredProducts)
}

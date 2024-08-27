package product

import "fmt"

/*
{
    "id" : 1,
    "name": "name"
    "description": "this is a cool product"
    "price": 10000
    "Category": "mobile"
}
*/

var (
	Products  = make([]Product, 0)
	latestPid = 1
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func init() {
	InsertDummyProduct()
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}

func InsertDummyProduct() {
	AddProduct("IPhone 15", "New iphone", "Mobile", 100000.00)
	AddProduct("Samsung", "New iphone", "Mobile", 10000.00)

	AddProduct("Micriwave Ovan", "Ovan", "Kitchen", 20000.00)

	AddProduct("Washing Machine", "Bosch", "Home", 40000.00)

	AddProduct("SomeProduct", "no desc", "no category", 2000.00)

}

func AddProduct(name, desc, cat string, price float64) {
	p := Product{Id: latestPid, Name: name, Desc: desc, Category: cat, Price: price}
	Products = append(Products, p)
	latestPid++
}

func GetProduct(pid int) (Product, error) {

	for _, p := range Products {
		if pid == p.Id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("invalid product id")
}

func UpdatePrice(pid int, price float64) error {

	for idx, p := range Products {
		if pid == p.Id {
			Products[idx].Price = price
			return nil
		}
	}
	return fmt.Errorf("invalid product id")
}

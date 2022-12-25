package models

type Product struct {
	Base
	Name         string  `json:"product_name"`
	Price        float64 `json:"product_price"`
	Image        []ProductImage
	Category     string  `json:"product_category"`
	SubCategory  string  `json:"product_subcategory"`
	Description  string  `json:"product_description"`
	Stock        bool    `json:"product_stock"`
	PricePerUnit float64 `json:"product_price_per_unit"`
	UnitsInStock float64 `json:"product_units_in_stock"`
	UnitsOnOrder float64 `json:"product_units_on_order"`
	CompanyId    string  `json:"product_company_id"`
}

func init() {
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.Create(&p)
	return p
}

func GetProduct() []Product {
	var product []Product
	db.Find(&product)
	return product
}

func DeleteProduct(id string) *Product {
	var product Product
	db.Where("Id =?", id).Delete(&product)
	return &product
}

func FindProductById(id string) *Product {
	var product Product
	db.Where("Id =?", id).Find(&product)
	return &product
}

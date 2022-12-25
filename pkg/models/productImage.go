package models

type ProductImage struct {
	Base
	Image     string `json:"product_image"`
	ProductId string `json:"product_id"`
}

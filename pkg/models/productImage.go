package models

type ProductImage struct {
	Base
	Image     string `json:"product_image"`
	Name      string `json:"product_name"`
	ProductId string `json:"product_id"`
}

func (p *ProductImage) CreateImage() *ProductImage {
	db.Model(&ProductImage{}).Create(p)
	return p
}

func GetProductImage(id string) []ProductImage {
	var images []ProductImage
	db.Model(&ProductImage{}).Where("product_id =?", id).Find(&images)
	return images
}

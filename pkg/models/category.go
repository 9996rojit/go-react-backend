package models

type Category struct {
	Base
	Name  string `json:"category_name"`
	Type  string `json:"category_type"`
	Image string `json:"category_image"`
}

func (c *Category) CreateCategory() *Category {
	db.Model(&Category{}).Create(c)
	return c
}

func GetCategory() []Category {
	var category []Category
	db.Model(&Category{}).Find(&category)
	return category
}

func GetCategoryById(id string) *Category {
	var category *Category
	db.Model(&Category{}).Where("Id = ?", id).Find(category)
	return category
}

func DeleteCategoryById(id string) *Category {
	var category *Category
	db.Model(&Category{}).Where("Id = ?", id).Delete(category)
	return category
}

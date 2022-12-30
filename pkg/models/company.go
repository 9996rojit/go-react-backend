package models

type Company struct {
	Base
	Name    string `json:"company_name"`
	Address string `json:"company_address"`
	Contact string `json:"company_contact"`
	Email   string `json:"company_email" gorm:"unique"`
	Type    string `json:"company_type"`
}

func (c *Company) CreateCompany() *Company {
	db.Model(&Company{}).Create(c)
	return c
}

func GetAllCompany() []Company {
	var companies []Company
	db.Model(&Company{}).Find(&companies)
	return companies

}

func GetCompanyByID(id string) *Company {
	var company Company
	db.Model(&Company{}).Where("id =?", id).First(&company)
	return &company
}

func DeleteCompanyByID(id string) *Company {
	var company Company
	db.Model(&Company{}).Where("id =?", id).Delete(&company)
	return &company
}

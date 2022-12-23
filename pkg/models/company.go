package models

type Company struct {
	Base
	Name    string `json:"company_name"`
	Address string `json:"company_address"`
	Contact string `json:"company_contact"`
	Email   string `json:"company_email"`
	Type    string `json:"company_type"`
}

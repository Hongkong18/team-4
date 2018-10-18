package models

type Institution struct {
	Name        string `json:"name"`
	Speciality  string `json:"speciality"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
}

type InstitutionInterface interface {
	ListAll() []*Institution
	GetById(id int64) *Institution
	Insert(Institution)
}

package data

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Org struct {
	ID               string `json:"id"`
	OrganisationName string `json:"organisationName"`
	OrgType          string `json:"orgType"`
	OrgSize          string `json:"orgSize"`
}

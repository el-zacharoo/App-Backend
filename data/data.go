package data

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Org struct {
	ComapanyName string `json:"companyName"`
}

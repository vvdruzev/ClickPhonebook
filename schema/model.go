package schema

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone []string
}

type Phone struct {
	Id          int
	PhoneNumber string
}

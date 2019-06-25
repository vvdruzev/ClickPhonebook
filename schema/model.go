package schema

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phones []string
}

type Phone struct {
	ContactId          int
	PhoneNumber string
}


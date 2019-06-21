package db

import (
	"ClicksPhonebook/schema"
)

type Repository interface {
	Close()
	AddContact(string, string) error
	List() (map[int]schema.Contact, map[int][]string, error)

}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}


func List() (map[int]schema.Contact, map[int][]string, error)  {
	return impl.List()
}

func AddContact(firstname string, lastname string) error  {
	return impl.AddContact(firstname,lastname)
}
package db

import (
	"ClickPhonebook/schema"
)

type Repository interface {
	Close()
	AddContact(string, string) error
	List() (map[int]schema.Contact, map[int][]string, error)
	AddPhone (int , string) error
	SelectItem (int) (map[int]schema.Contact, map[int][]string, error)
	Delete (int) error
	Update (schema.Contact,[]string)  error
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

func AddPhone(id int, number string) error  {
	return impl.AddPhone(id,number)
}

func SelectItem( id int) (map[int]schema.Contact, map[int][]string, error)  {
	return impl.SelectItem(id)
}

func Delete(id int) error {
	return impl.Delete(id)
}

func Update(contact schema.Contact,phones []string) error  {
	return impl.Update(contact,phones)
}
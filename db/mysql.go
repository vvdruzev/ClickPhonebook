package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"ClickPhonebook/schema"
	"ClickPhonebook/logger"
)

type Mysqlrepo struct {
	Db *sql.DB
}

func NewMysqlrepo(dsn *string) (*Mysqlrepo, error) {
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		return nil, err
	}
	return &Mysqlrepo{
		db,
	}, nil
}

type dbError struct {
	method string
	Err    error
}

func (re *dbError) Error() string {
	return fmt.Sprintf(
		"DB error:  %s, err: %v",
		re.method,
		re.Err,
	)
}

func (db Mysqlrepo) Close() {
	db.Db.Close()
}

func (db Mysqlrepo) AddContact(firstname string, lastname string) error {
	result, err := db.Db.Exec(
		"INSERT INTO Contacts (firstname, lastname) VALUES (?, ?)",
		firstname,
		lastname,
	)
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}

	lastID, err := result.LastInsertId()

	logger.Info("Insert - RowsAffected", affected, "LastInsertId: ", lastID)

	return nil
}

func (db Mysqlrepo) AddPhone(idContact int, phone string) error {
	result, err := db.Db.Exec(
		"INSERT INTO Phonenumber (id,phonenumber) VALUES (?,?)",
		idContact,
		phone,
	)
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}

	lastID, err := result.LastInsertId()

	logger.Info("Insert - RowsAffected", affected, "LastInsertId: ", lastID)

	return nil
}

func (db Mysqlrepo) List() (map[int]schema.Contact, map[int][]string, error) {
	contacts := make(map[int]schema.Contact)
	phones := make(map[int][]string)
	sqlStr := "select id, firstname, lastname from Contacts"
	rows, err := db.Db.Query(sqlStr)
	for rows.Next() {
		contact := &schema.Contact{}
		err = rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName)
		if err != nil {
			logger.Error("Can't select rows", err)
			return nil, nil, err
		}
		contacts[contact.Id] = *contact
	}

	rows, err = db.Db.Query("select id, Phonenumber from Phonenumber")
	for rows.Next() {
		phone := &schema.Phone{}
		err = rows.Scan(&phone.Id, &phone.PhoneNumber)
		if err != nil {
			logger.Error("Can't select rows", err)
			return nil, nil, err
		}
		phones[phone.Id] = append(phones[phone.Id],phone.PhoneNumber)
	}
	rows.Close()
	return contacts, phones, nil
}

func (db Mysqlrepo) Delete(id int) error {
	result, err := db.Db.Exec(
		"DELETE FROM Contacts WHERE id = ?",
		id,
	)
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}

	logger.Info("Delete - RowsAffected", affected)

	result, err = db.Db.Exec(
		"DELETE FROM Phonenumber WHERE id = ?",
		id,
	)
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}
	affected, err = result.RowsAffected()
	if err != nil {
		logger.Error("Can't add rows", err)
		return err
	}

	logger.Info("Delete - RowsAffected", affected)

	return nil

}

func (db Mysqlrepo) Update (contact schema.Contact, phones []schema.Phone)  {

}
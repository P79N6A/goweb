package models

import (
	"goweb/databases"
	. "goweb/utils"
)

//定义person类型结构
type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var sqldb = databases.SqlDB
func (p *Person) AddPerson() (id int64, err error) {
	rs, err := sqldb.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := sqldb.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *Person) GetPerson() (person Person, err error) {
	err = sqldb.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", p.Id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	stmt, err := sqldb.Prepare("UPDATE person SET first_name=?, last_name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := sqldb.Exec("DELETE FROM person WHERE id=?", p.Id)
	CheckErr(err)
	ra, err = rs.RowsAffected()
	return
}

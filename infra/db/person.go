package db

import (
	"database/sql"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
)

func InsertPerson(db *sql.DB, p *entity.Person) {
	stmp, _ := db.Prepare("INSERT INTO person(id,name,age) VALUES(?,?,?)")
	stmp.Exec(p.Id, p.Name, p.Age)
	defer stmp.Close()
}
func UpdatePerson(db *sql.DB, p *entity.Person) error {
	stmp, err := db.Prepare("UPDATE person SET name=?,age=? WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.Name, p.Age, p.Id)
	defer stmp.Close()
	return nil
}

func DeletePerson(db *sql.DB, p *entity.Person) error {
	stmp, err := db.Prepare("DELETE FROM person WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.Id)
	defer stmp.Close()
	return nil
}

func GetPerson(db *sql.DB, id string) (*entity.Person, error) {
	stmp, err := db.Prepare("SELECT id,name,age FROM person WHERE id=?")
	if err != nil {
		return nil, err
	}
	var p entity.Person
	stmp.QueryRow(id).Scan(&p.Id, &p.Name, &p.Age)
	defer stmp.Close()
	return &p, nil
}

func GetAllPerson(db *sql.DB) ([]*entity.Person, error) {
	rows, err := db.Query("SELECT id,name,age FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var persons []*entity.Person

	for rows.Next() {
		var p entity.Person
		err = rows.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &p)
	}

	return persons, nil
}

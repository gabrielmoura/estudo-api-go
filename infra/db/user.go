package db

import (
	"database/sql"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
)

func GetAllUser(db *sql.DB) ([]*entity.User, error) {
	rows, err := db.Query("SELECT id,name,email,password FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*entity.User

	for rows.Next() {
		var p entity.User
		err = rows.Scan(&p.ID, &p.Name, &p.Email, &p.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &p)
	}

	return users, nil
}

func InsertUser(db *sql.DB, u *entity.User) (bool, error) {
	stmp, err := db.Prepare("INSERT INTO user(id,name,email,password) VALUES(?,?,?,?)")
	if err != nil {
		return false, err
	}
	stmp.Exec(u.ID, u.Name, u.Email, u.Password)

	defer stmp.Close()

	return true, nil
}

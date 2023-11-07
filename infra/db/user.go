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

func GetOneUser(db *sql.DB, id string) (*entity.User, error) {
	stmp, err := db.Prepare("SELECT id,name,email,password FROM user WHERE id=?")
	if err != nil {
		return nil, err
	}
	var p entity.User
	stmp.QueryRow(id).Scan(&p.ID, &p.Name, &p.Email, &p.Password)
	defer stmp.Close()
	return &p, nil
}

func InsertUser(db *sql.DB, u *entity.User) (bool, error) {
	stmp, err := db.Prepare("INSERT INTO user(id,name,email,password,created_at) VALUES(?,?,?,?,?)")
	if err != nil {
		return false, err
	}
	stmp.Exec(u.ID, u.Name, u.Email, u.Password, u.CreatedAt)

	defer stmp.Close()

	return true, nil
}
func UpdateUser(db *sql.DB, p *entity.User) error {
	stmp, err := db.Prepare("UPDATE user SET name=?,email=? WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.Name, p.Email, p.ID)
	defer stmp.Close()
	return nil
}

func DeleteUser(db *sql.DB, p *entity.User) error {
	stmp, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.ID)
	defer stmp.Close()
	return nil
}

package db

import (
	"database/sql"

	"github.com/gabrielmoura/estudo-api-go/internal/entity"
)

func GetAllCart(db *sql.DB) ([]*entity.Cart, error) {
	rows, err := db.Query("SELECT id,products FROM cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var carts []*entity.Cart

	for rows.Next() {
		var p entity.Cart
		err = rows.Scan(&p.ID, &p.Product)
		if err != nil {
			return nil, err
		}
		carts = append(carts, &p)
	}

	return carts, nil
}

func GetOneCart(db *sql.DB, id string) (*entity.Cart, error) {
	stmp, err := db.Prepare("SELECT id,products FROM cart WHERE id=?")
	if err != nil {
		return nil, err
	}
	var p entity.Cart
	stmp.QueryRow(id).Scan(&p.ID, &p.Product)
	defer stmp.Close()
	return &p, nil
}

func InsertCart(db *sql.DB, p *entity.Cart) (bool, error) {
	stmp, err := db.Prepare("INSERT INTO cart(id,products) VALUES(?,?)")
	if err != nil {
		return false, err
	}
	stmp.Exec(p.ID, p.Product)

	defer stmp.Close()

	return true, nil
}

func DeleteCart(db *sql.DB, p *entity.Cart) error {
	stmp, err := db.Prepare("DELETE FROM cart WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.ID)
	defer stmp.Close()
	return nil
}

func UpdateCart(db *sql.DB, p *entity.Cart) error {
	stmp, err := db.Prepare("UPDATE cart SET product=? WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.Product, p.ID)
	defer stmp.Close()
	return nil
}

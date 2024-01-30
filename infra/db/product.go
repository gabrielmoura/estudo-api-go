package db

import (
	"database/sql"
	"github.com/gabrielmoura/estudo-api-go/internal/util"
	"time"

	"github.com/gabrielmoura/estudo-api-go/internal/entity"
)

func GetAllProduct(db *sql.DB) ([]*entity.Product, error) {
	rows, err := db.Query("SELECT id,name,price,created_at FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product

	for rows.Next() {
		var p entity.Product
		var createdAtString string
		if err = rows.Scan(&p.ID, &p.Name, &p.Price, &createdAtString); err != nil {
			return nil, err
		}

		if p.CreatedAt, err = util.ConvertStringToTime(createdAtString); err != nil {
			return nil, err
		}

		products = append(products, &p)
	}

	return products, nil
}

func GetOneProduct(db *sql.DB, id string) (*entity.Product, error) {
	stmp, err := db.Prepare("SELECT id,name,price,created_at FROM product WHERE id=?")
	if err != nil {
		return nil, err
	}
	var p entity.Product
	var createdAtString string

	if err := stmp.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price, &createdAtString); err != nil {
		return nil, err
	}
	if p.CreatedAt, err = util.ConvertStringToTime(createdAtString); err != nil {
		return nil, err
	}

	defer stmp.Close()
	return &p, nil
}

func InsertProduct(db *sql.DB, p *entity.Product) (bool, error) {
	stmp, err := db.Prepare("INSERT INTO product(id,name,price,created_at) VALUES(?,?,?,?)")
	if err != nil {
		return false, err
	}
	stmp.Exec(p.ID, p.Name, p.Price, p.CreatedAt.Format(time.RFC3339))

	defer stmp.Close()

	return true, nil
}

func DeleteProduct(db *sql.DB, p *entity.Product) error {
	stmp, err := db.Prepare("DELETE FROM product WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.ID)
	defer stmp.Close()
	return nil
}

func UpdateProduct(db *sql.DB, p *entity.Product) error {
	stmp, err := db.Prepare("UPDATE product SET name=?,price=? WHERE id=?")
	if err != nil {
		return err
	}
	stmp.Exec(p.Name, p.Price, p.ID)
	defer stmp.Close()
	return nil
}

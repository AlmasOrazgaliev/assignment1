package controller

import (
	"database/sql"
	"github.com/AlmasOrazgaliev/assignment1/model"
)

type Controller struct {
	DB *sql.DB
}

func New(db *sql.DB) *Controller {
	return &Controller{
		DB: db,
	}
}

func (c *Controller) CreateUser(u *model.User) error {
	//err := u.BeforeCreate()
	//if err != nil {
	//	return err
	//}
	return c.DB.QueryRow(
		"INSERT INTO users (email,password,is_seller) VALUES ($1,$2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
		u.IsSeller
	).Scan(&u.Id)
}

func (c *Controller) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := c.DB.QueryRow(
		"SELECT id, email, password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (c *Controller) CreateSeller(u *model.User) error {

}
func (c *Controller) AllItems() []model.Item {
	var items []model.Item
	rows, err := c.DB.Query("SELECT id,name,price,description FROM items")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.Description)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}

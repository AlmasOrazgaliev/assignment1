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
	_, err := c.DB.Exec(
		"INSERT INTO users (email,password,is_seller) VALUES ($1,$2)",
		u.Email,
		u.EncryptedPassword,
		u.IsSeller,
	)
	return err
}

func (c *Controller) FindUser(email string, password string) (*model.User, error) {
	u := &model.User{}
	if err := c.DB.QueryRow(
		"SELECT id, email, password FROM users WHERE email=$1 AND password=$2",
		email,
		password,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (c *Controller) CreateItem(item *model.Item) error {
	_, err := c.DB.Exec("INSERT INTO items (name,price,description) VALUES ($1,$2,$3)",
		item.Name,
		item.Price,
		item.Description)
	return err
}
func (c *Controller) AllItems() []model.Item {
	var items []model.Item
	items = []model.Item{}
	rows, err := c.DB.Query("SELECT id,name,price,description, sold, rating FROM items")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.Description, &item.Sold, &item.Rating)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return items
}

func (c *Controller) SearchByPrice(min int, max int) []model.Item {
	items := c.AllItems()
	var sorted []model.Item
	for _, item := range items {
		if item.Rating/item.Sold >= min && item.Rating/item.Sold <= max {

		}
	}
}

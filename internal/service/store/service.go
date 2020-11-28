package store

import (
	"fmt"
	"github.com/SoleMer/dulceCaliGo/internal/config"
	"github.com/jmoiron/sqlx"
)

//ClothingItem
type ClothingItem struct {
	ID   int64
	Name string
}

func NewClothingItem(n string) (ClothingItem) {
	return ClothingItem{0, n}
}

// Service ...
type Service interface {
	AddItem(item ClothingItem) error
	FindByID(int) *ClothingItem
	FindAll() []*ClothingItem
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New...
func New(db *sqlx.DB, c *config.Config) (service, error) {
	return service{db, c}, nil
}

func (s service) AddItem(n string) error {

	insertItem  := `INSERT INTO clothes (name) VALUES (?)`
	m := fmt.Sprintf(n)
	s.db.MustExec(insertItem, m)

	return nil
}

func (s service) FindById(id int) (*ClothingItem, error) {
	getItem := `SELECT * FROM clothes WHERE id=?;`
	s.db.MustExec(getItem, id)

	var it ClothingItem
	err := s.db.QueryRowx(getItem, id).StructScan(&it)

	if err != nil {
		return nil, err
	}

	return &it, nil

}

func (s service) FindAll() []*ClothingItem {
	var list []*ClothingItem
	s.db.Select(&list, "SELECT * FROM clothes")
	return list
}

func (s service) DeleteItem(id int) error {
	dltItem := `DELETE FROM clothes WHERE id=?;`
	_, err := s.db.Exec(dltItem, id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) EditItem(id int, t string) (*ClothingItem, error) {
	putItem := `UPDATE clothes SET name=? WHERE id=?;`
	s.db.MustExec(putItem, t, id)

	result, err := s.FindById(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}
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
	Price int
	Stock int
}

func NewClothingItem(n string, p int, s int) (ClothingItem) {
	return ClothingItem{0, n, p, s}
}

// Service ...
type Service interface {
	AddItem(item ClothingItem) error
	FindByID(int) *ClothingItem
	FindAll() []*ClothingItem
	//TODO agregar funcs
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New...
func New(db *sqlx.DB, c *config.Config) (service, error) {
	return service{db, c}, nil
}

func (s service) AddItem(n string, p int, q int) (int64, error) {

	insertItem  := `INSERT INTO clothes (name, price, stock) VALUES (?, ?, ?)`

	var cItem = ClothingItem{
		Name: n,
		Price: p,
		Stock: q}

	fmt.Println("service")
	fmt.Println(cItem.Name)
	fmt.Println(cItem.Price)
	fmt.Println(cItem.Stock)

	result := s.db.MustExec(insertItem, cItem.Name, cItem.Price, cItem.Stock)

	lastID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (s service) FindById(id int64) (*ClothingItem, error) {
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

func (s service) EditItem(id int64, n string, p int, q int) (*ClothingItem, error) {
	putItem := `UPDATE clothes SET name=?, price=?, stock=? WHERE id=?;`

	var cItem = ClothingItem{
		Name: n,
		Price: p,
		Stock: q,
	}

	s.db.MustExec(putItem, cItem.Name, cItem.Price, cItem.Stock, id)


	result, err := s.FindById(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}
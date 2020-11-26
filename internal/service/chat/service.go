package chat

import (
	"fmt"
	"github.com/SoleMer/proyectoGO/internal/config"
	"github.com/jmoiron/sqlx"
)

//Message ...
type Message struct {
	ID   int64
	Text string
}

func NewMessage(t string) (Message) {
	return Message{0, t}
}

// Service ...
type Service interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New...
func New(db *sqlx.DB, c *config.Config) (service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(t string) error {

	insertMessage  := `INSERT INTO messages (text) VALUES (?)`
	m := fmt.Sprintf(t)
	s.db.MustExec(insertMessage, m)

	return nil
}

func (s service) FindById(id int64) (*Message, error) {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	s.db.Select(&list, "SELECT * FROM messages")
	return list
}

func (s service) DeleteMsg(id int) error {
	return nil
}

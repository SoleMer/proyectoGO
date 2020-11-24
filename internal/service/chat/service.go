package chat

import (
	"github.com/SoleMer/proyectoGO/internal/config"
	"github.com/jmoiron/sqlx"
)

//Message ...
type Message struct {
	ID   int64
	Text string
}

//ChatService ...
type ChatService interface {
	AddMessage(Message) error
	FindById(int) *Message
	FindAll() []*Message
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New...
func New(db *sqlx.DB, c *config.Config) (ChatService, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(m Message) error {
	return nil
}

func (sservice) FindById(ID int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	s.db.Select(&list, "SELECT * FROM messages")
	return list
}

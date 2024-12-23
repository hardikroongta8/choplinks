package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hardikroongta8/choplinks/model"
)

type Store interface {
	CreateUser(*model.User) error
	GetUserByID(string) (*model.User, error)
	CreateURLMap(*model.URLMap) error
	DeleteURLMapByID(string) error
	GetAllURLMapsByUserID(string) ([]*model.URLMap, error)
	GetURLMapByID(string) (*model.URLMap, error)
}

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(uri string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLStore{db: db}, nil
}

func (s *MySQLStore) Init() error {
	err := s.createUserTable()
	if err != nil {
		return err
	}
	err = s.createURLMapsTable()
	if err != nil {
		return err
	}
	return nil
}

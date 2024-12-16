package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Storage interface {
	CreateUser(*User) error
	GetUserByID(string) (*User, error)
	CreateURLMap(*URLMap) error
	DeleteURLMapUsingShortenedURL(string) error
	GetAllURLMapsByUserID(string) ([]URLMap, error)
	GetURLMapByShortenedURL(string) (*URLMap, error)
}

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() (*MySQLStore, error) {
	cfg := GetConfig()
	db, err := sql.Open("mysql", cfg.DB.URI)
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

func (s *MySQLStore) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS user (
		id CHAR(36) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	log.Println("Successfully created User Table!")
	return nil
}

func (s *MySQLStore) createURLMapsTable() error {
	query := `CREATE TABLE IF NOT EXISTS url_maps (
		shortened_url_path VARCHAR(255) PRIMARY KEY,
		original_url VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		user_id CHAR(36) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(id)
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *MySQLStore) CreateUser(u *User) error {
	query := fmt.Sprintf(`INSERT INTO user(id, name, email)
	VALUES ('%s', '%s', '%s')`, u.ID, u.Name, u.Email)

	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) GetUserByID(id string) (*User, error) {
	return nil, nil
}
func (s *MySQLStore) CreateURLMap(u *URLMap) error {
	query := fmt.Sprintf(`INSERT INTO url_maps
    (shortened_url_path, original_url, user_id)
	VALUES ('%s', '%s', '%s')`, u.ShortenedURLPath, u.OriginalURL, u.UserID)

	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) DeleteURLMapUsingShortenedURL(url string) error {
	return nil
}
func (s *MySQLStore) GetAllURLMapsByUserID(id string) ([]URLMap, error) {
	rows, err := s.db.Query("SELECT * FROM url_maps WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlMaps []URLMap

	for rows.Next() {
		var urlMap URLMap
		if err := rows.Scan(
			&urlMap.ShortenedURLPath,
			&urlMap.OriginalURL,
			&urlMap.CreatedAt,
			&urlMap.UpdatedAt,
			&urlMap.UserID,
		); err != nil {
			return urlMaps, err
		}
		urlMaps = append(urlMaps, urlMap)
	}

	err = rows.Err()
	return urlMaps, err
}
func (s *MySQLStore) GetURLMapByShortenedURL(url string) (*URLMap, error) {
	return nil, nil
}

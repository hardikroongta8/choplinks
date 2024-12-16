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
	DeleteURLMapByID(string) error
	GetAllURLMapsByUserID(string) ([]*URLMap, error)
	GetURLMapByID(string) (*URLMap, error)
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
	query := `CREATE TABLE IF NOT EXISTS users (
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
		id VARCHAR(255) PRIMARY KEY,
		original_url VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		user_id CHAR(36) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	log.Println("Successfully created URLMaps Table!")
	return nil
}

func (s *MySQLStore) CreateUser(u *User) error {
	query := fmt.Sprintf(`INSERT INTO users(id, name, email)
	VALUES ('%s', '%s', '%s')`, u.ID, u.Name, u.Email)

	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) GetUserByID(id string) (*User, error) {
	query := fmt.Sprintf(`SELECT * FROM users
	WHERE id = '%s'`, id)
	row := s.db.QueryRow(query)
	user := new(User)
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	); err != nil {
		return nil, err
	}
	return user, nil
}
func (s *MySQLStore) CreateURLMap(u *URLMap) error {
	query := fmt.Sprintf(`INSERT INTO url_maps
    (id, original_url, user_id)
	VALUES ('%s', '%s', '%s')`, u.ID, u.OriginalURL, u.UserID)

	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) DeleteURLMapByID(id string) error {
	query := fmt.Sprintf(`DELETE FROM url_maps
	WHERE id = '%s'`, id)
	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) GetAllURLMapsByUserID(id string) ([]*URLMap, error) {
	rows, err := s.db.Query("SELECT * FROM url_maps WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlMaps []*URLMap

	for rows.Next() {
		urlMap := new(URLMap)
		if err := rows.Scan(
			&urlMap.ID,
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
func (s *MySQLStore) GetURLMapByID(path string) (*URLMap, error) {
	query := fmt.Sprintf(`SELECT * FROM url_maps
	WHERE id = '%s'`, path)
	row := s.db.QueryRow(query)
	urlMap := new(URLMap)
	if err := row.Scan(
		&urlMap.ID,
		&urlMap.OriginalURL,
		&urlMap.CreatedAt,
		&urlMap.UpdatedAt,
		&urlMap.UserID,
	); err != nil {
		return nil, err
	}
	return urlMap, nil
}

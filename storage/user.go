package storage

import (
	"fmt"
	"github.com/hardikroongta8/choplinks/model"
	"log"
)

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

func (s *MySQLStore) CreateUser(u *model.User) error {
	query := fmt.Sprintf(`INSERT INTO users(id, name, email)
	VALUES ('%s', '%s', '%s')`, u.ID, u.Name, u.Email)

	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) GetUserByID(id string) (*model.User, error) {
	query := fmt.Sprintf(`SELECT * FROM users
	WHERE id = '%s'`, id)
	row := s.db.QueryRow(query)
	user := new(model.User)
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	); err != nil {
		return nil, err
	}
	return user, nil
}

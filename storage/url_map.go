package storage

import (
	"fmt"
	"github.com/hardikroongta8/choplinks/model"
	"log"
)

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

func (s *MySQLStore) CreateURLMap(u *model.URLMap) error {
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
func (s *MySQLStore) GetAllURLMapsByUserID(id string) ([]*model.URLMap, error) {
	rows, err := s.db.Query("SELECT * FROM url_maps WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlMaps []*model.URLMap

	for rows.Next() {
		urlMap := new(model.URLMap)
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
func (s *MySQLStore) GetURLMapByID(path string) (*model.URLMap, error) {
	query := fmt.Sprintf(`SELECT * FROM url_maps
	WHERE id = '%s'`, path)
	row := s.db.QueryRow(query)
	urlMap := new(model.URLMap)
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

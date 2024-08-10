package postgres

import "log"

func (s *Store) CreateToken(phone, token string) error {
	row, err := s.db.Exec("INSERT INTO users(token) VALUES ($1) WHERE phone = $2", token, phone)
	if err != nil {
		return err
	}

	affectedRow, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow != 1 {
		log.Printf("affected more rows than 1: %d", affectedRow)
	}

	return err
}

func (s *Store) GetToken(phone string) (string, error) {
	row, err := s.db.Queryx("SELECT token FROM users WHERE phone = $1", phone)
	if err != nil {
		return "", err
	}

	var token string
	for row.Next() {
		err = row.Scan(&token)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

package postgres

import (
	"auth-service/internal/model"
	"github.com/google/uuid"
	"log"
)

func (s *Store) SaveUser(signModel model.SignUp) error {
	query := `INSERT INTO(iin, name, surname, phone, password) users VALUES(:=iin, :=name, :=surname, :=phone, :=password)`
	row, err := s.db.NamedExec(query, signModel)
	if err != nil {
		return err
	}

	affectedRow, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow != 1 {
		log.Printf("affected more rows than need: %d", affectedRow)
	}

	err = s.insertToken(signModel.Phone)
	if err != nil {
		return err
	}

	return err
}

func (s *Store) insertToken(phone string) error {
	token := uuid.New().String()

	row, err := s.db.Exec("INSERT INTO users(token) VALUES($1) WHERE phone=$2", token, phone)
	if err != nil {
		return err
	}

	affectedRows, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows != 1 {
		log.Printf("affected more than 1 row: %d", affectedRows)
	}

	return nil
}

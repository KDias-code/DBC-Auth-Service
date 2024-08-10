package databases

import "github.com/jmoiron/sqlx"

func PostgresConnection(dbStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

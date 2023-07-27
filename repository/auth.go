package repository

import "github.com/tckthecreator/clean_arch_go/model"

func (db *repository) CreateAuth(auth *model.Auth) (int64, error) {
	sqmt, err := db.pg.Prepare(`INSERT INTO auth("username", "email", "password", "country", "refresh_token") VALUES ($1, $2, $3, $4, $5) RETURNING id`)
	if err != nil {
		return 0, err
	}

	row := sqmt.QueryRow(auth.Username, auth.Email, auth.Password, auth.Country, auth.RefreshToken)
	if row == nil {
		return 0, err
	}

	var id int64

	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

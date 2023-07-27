package repository

import (
	"database/sql"

	"github.com/tckthecreator/clean_arch_go/model"
)

type repository struct {
	pg *sql.DB
}

func NewRepository(pg *sql.DB) model.Repository {
	return &repository{
		pg,
	}
}

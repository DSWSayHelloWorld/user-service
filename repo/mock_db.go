package repo

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

func GetUserByIDMock() *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
		AddRow(123, "Virat", "Kohli", "virat@example.com")
	query := "SELECT * FROM user where id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(0).WillReturnRows(rows)
	return db
}

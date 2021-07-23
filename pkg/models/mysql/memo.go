package mysql

import (
	"database/sql"

	"github.com/zemags/go_workshop_4/pkg/models"
)

type MemoModel struct {
	DB *sql.DB
}

func (m *MemoModel) Insert(title, content, expires string) (int, error) {

	query := `insert into memo (title, content, created, expires)
		values (?, ?, utc_timestamp, date_add(utc_timestamp(), interval ? day))`

	result, err := m.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m *MemoModel) Get(id int) (*models.Memo, error) {
	return nil, nil
}

func (m *MemoModel) Latest() ([]*models.Memo, error) {
	return nil, nil
}

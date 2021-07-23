package mysql

import (
	"database/sql"
	"errors"

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
	query := `select id, title, content, created, expires from memo
	where expires > utc_timestamp() and id = ?`

	row := m.DB.QueryRow(query, id)

	mm := &models.Memo{}

	if err := row.Scan(&mm.ID, &mm.Title, &mm.Content, &mm.Created, &mm.Expires); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return mm, nil
}

func (m *MemoModel) Latest() ([]*models.Memo, error) {
	return nil, nil
}

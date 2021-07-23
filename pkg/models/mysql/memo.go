package mysql

import (
	"database/sql"

	"github.com/zemags/go_workshop_4/pkg/models"
)

type MemoModel struct {
	DB *sql.DB
}

func (m *MemoModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *MemoModel) Get(id int) (*models.Memo, error) {
	return nil, nil
}

func (m *MemoModel) Latest() ([]*models.Memo, error) {
	return nil, nil
}

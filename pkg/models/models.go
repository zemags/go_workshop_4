package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: cannot find record")

type Memo struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

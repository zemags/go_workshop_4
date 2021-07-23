package models

import (
	"errors"
	"time"
)

// like exceptions in python
var ErrNoRecord = errors.New("models: cannot find record")

type Memo struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

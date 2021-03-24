package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

type Task struct {
	gorm.Model
	Title    string
	Status   string
	Due_date time.Time
	UserID   uint
}

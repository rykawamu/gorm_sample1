package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

type User struct {
	gorm.Model // gorm.Model 構造体（ID,CreatedAt,UpdatedAt,DeletedAt）の追加
	Name       string
	Email      string
	Tasks      []Task
}

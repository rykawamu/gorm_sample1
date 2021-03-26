package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

type State struct {
	gorm.Model
	StateId   int    `json:"stateid"`
	StateName string `json:"statename"`
}

package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

type User struct {
	gorm.Model        // gorm.Model 構造体（ID,CreatedAt,UpdatedAt,DeletedAt）の追加
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Tasks      []Task
}

func FindUser(u *User) User {
	var user User
	mydb.Where(u).First(&user)
	return user
}

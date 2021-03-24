package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

type Task struct {
	gorm.Model
	Title    string `json:"title"`
	Status   string `json:"status"`
	Due_date time.Time
	UserID   uint
}

type Tasks []Task

func FindTasks(t *Task) Tasks {
	var tasks Tasks
	mydb.Where(t).Find(&tasks)
	return tasks
}

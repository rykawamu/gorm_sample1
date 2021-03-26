package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

type Task struct {
	gorm.Model
	Title    string `json:"title"`
	StateId  uint   `json:"stateid"`
	Due_date time.Time
	UserID   uint
}

type Tasks []Task

type TaskResult struct {
	Title     string `json:"title"`
	StateId   uint   `json:"stateid"`
	StateName string `json:"statename"`
}

type TaskResults []TaskResult

func FindTasks(t *Task) Tasks {
	var tasks Tasks
	mydb.Where(t).Find(&tasks)
	return tasks
}

func FindTaskResults(t *Task) TaskResults {
	var taskresults TaskResults
	mydb.Model(&Tasks{}).
		Select("tasks.title, tasks.state_id, states.state_name").
		Joins("left join states on tasks.state_id = states.state_id").
		Where("user_id = ?", t.UserID).
		Scan(&taskresults)
	return taskresults
}

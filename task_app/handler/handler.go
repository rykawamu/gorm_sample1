package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	mymodel "taskapp/model"
)

func GetTasks(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := mymodel.FindUser(&mymodel.User{Model: gorm.Model{ID: uint(uid)}}); user.ID == 0 {
		//fmt.Println("***************")
		return echo.ErrNotFound
	}

	tasks := mymodel.FindTasks(&mymodel.Task{UserID: uint(uid)})
	return c.JSON(http.StatusOK, tasks)
}

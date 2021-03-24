package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //ブランク識別子: 依存関係を解決するためのimport
)

const LocalDB string = "./db/test.db"

type MyModel struct {
	Dialector string
	DbKind    string
	Db        *gorm.DB
}

func (m *MyModel) connect_database() (db *gorm.DB, err error) {
	m.Dialector = "sqlite3"
	m.DbKind = "./db/test.db"

	db, err = gorm.Open(m.Dialector, m.DbKind)
	if err != nil {
		panic("failed to connect database")
	}
	m.Db = db
	return db, err
}

func (m *MyModel) Init() {
	// connect database
	db, err := m.connect_database()
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Task{})

	// Delete
	//db.Where("1 = 1").Delete(&Task{}) // gorm.DeletedAtがある場合、soft deleteになる
	db.Exec("DELETE FROM tasks") // レコード自体を消したいなら、こっち
	db.Exec("DELETE FROM users")
}

func (m *MyModel) InsertInitData() {
	db, err := m.connect_database()
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	db.Create(&User{Name: "Alice", Email: "alice@example.com",
		Tasks: []Task{{Title: "work1", Status: "start", Due_date: time.Now()},
			{Title: "work2", Status: "stop", Due_date: time.Now()},
		},
	})
	db.Create(&User{Name: "Betty", Email: "Betty@example.com",
		Tasks: []Task{{Title: "work3", Status: "start", Due_date: time.Now()},
			{Title: "work1", Status: "Cancel", Due_date: time.Now()},
		},
	})
	db.Create(&User{Name: "Carmichael", Email: "Carmichael@example.com",
		Tasks: []Task{{Title: "work5", Status: "start", Due_date: time.Now()},
			{Title: "work6", Status: "Cancel", Due_date: time.Now()},
			{Title: "work7", Status: "Cancel", Due_date: time.Now()},
		},
	})
	db.Create(&User{Name: "George", Email: "George@example.com", Tasks: []Task{}})

}

// 初期データのインサート情報のデバック確認用
func (m *MyModel) InitDataPrint() {
	db, err := m.connect_database()
	if err != nil {
		panic("failed to connect database")
	}

	// Read
	var user User
	// first matched record
	db.First(&user)
	fmt.Printf("check1: %v\n", &user)
	fmt.Println("-----")

	// Get first matched record
	var where_user User
	db.Where("name = ?", "Carmichael").First(&where_user)
	fmt.Printf("check2: %v\n", &where_user)
	fmt.Println("-----")

	// Get all records
	var users []User
	db.Find(&users)
	fmt.Printf("check3: %v\n", &users)
}

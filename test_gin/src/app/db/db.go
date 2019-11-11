package db

import (
	"github.com/jinzhu/gorm"
	// GORM が使用するため、main.go 上では使わなくても宣言する必要あり
	//（なので使ってないよという意味でアンダースコア）
	_ "github.com/mattn/go-sqlite3"
)

// Connect : DB 接続
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")

	if err != nil {
		panic("can not open database!")
	}

	return db
}

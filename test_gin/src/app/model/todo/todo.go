package todo

import (
	// ローカルパッケージパス
	"app/db"
	"app/enum"

	"github.com/jinzhu/gorm"
	// GORM が使用するため、main.go 上では使わなくても宣言する必要あり
	//（なので使ってないよという意味でアンダースコア）
	_ "github.com/mattn/go-sqlite3"
)

// Todo : TODO 情報を管理する構造体
type Todo struct {
	gorm.Model
	Text   string
	Status enum.Status
}

// Migrate : DB 初期化
func Migrate() {
	db := db.Connect()
	defer db.Close()

	db.AutoMigrate(&Todo{})
}

// Insert : データ追加
func Insert(text string, status enum.Status) {
	db := db.Connect()
	defer db.Close()

	db.Create(&Todo{Text: text, Status: status})
}

// Update : データ更新
func Update(id int, text string, status enum.Status) {
	db := db.Connect()
	defer db.Close()

	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
}

// Delete : データ削除
func Delete(id int) {
	db := db.Connect()
	defer db.Close()

	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
}

// All : データ全取得
func All() []Todo {
	db := db.Connect()
	defer db.Close()

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	return todos
}

// FindByID : データ 1レコード取得
func FindByID(id int) Todo {
	db := db.Connect()
	defer db.Close()

	var todo Todo
	db.First(&todo, id)
	return todo
}

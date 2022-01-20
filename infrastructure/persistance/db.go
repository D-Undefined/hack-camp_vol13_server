package persistance

import (
	"fmt"
	"os"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SqlHandler struct {
	db *gorm.DB
}

// db接続とモデルのmigrate
func NewDB() *SqlHandler {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@db:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		"5432",
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Thread{}, &model.Comment{})

	sqlhandler := new(SqlHandler)
	sqlhandler.db = db
	return sqlhandler
}

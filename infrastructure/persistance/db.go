package persistance

import (
	"fmt"
	"os"
	"time"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SqlHandler struct {
	db *gorm.DB
}

// db接続とモデルのmigrate
func NewDB() *SqlHandler {
	var connectionString string

	if os.Getenv("APP_MODE") == "production" {
		// 本番環境
		connectionString = os.Getenv("DATABASE_URL")
	} else {
		// 開発環境
		connectionString = fmt.Sprintf(
			"postgres://%s:%s@db:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			"5432",
			os.Getenv("POSTGRES_DB"),
		)
	}

	// 稼働待ち
	for i := 0; i < 10; i++ {
		_, err := gorm.Open("postgres", connectionString)
		if err == nil {
			fmt.Printf(os.Getenv("DATABASE_URL") + "### connect.\n")
			break
		}
		fmt.Printf("### failed to connect database. connect again.\n")
		time.Sleep(3 * time.Second)
	}

	db, _ := gorm.Open("postgres", connectionString)

	db.AutoMigrate(&model.User{}, &model.Thread{}, &model.Comment{})
	db.AutoMigrate(&model.VoteComment{}, &model.VoteThread{})

	sqlhandler := new(SqlHandler)
	sqlhandler.db = db

	return sqlhandler
}

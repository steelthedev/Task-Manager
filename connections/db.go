package model

import (
	"log"

	"github.com/steelthedev/task-go/pkg/task/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Task{})

	return db
}

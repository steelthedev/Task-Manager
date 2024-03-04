package connections

import (
	"log"

	"github.com/steelthedev/task-go/pkg/task/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(dbUrl string) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatalln("could not connect o jare")
	}

	db.AutoMigrate(&model.Task{})

}

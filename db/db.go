package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Goal struct {
	gorm.Model
	Name        string
	Description string
	ParentID int32
	Parent *Goal `gorm:"foreignkey:ParentID"`
}

type DB struct {
	db *gorm.DB
}

func Connect() DB {
	db, err := gorm.Open(mysql.Open(
		"leighton:123456@tcp(127.0.0.1:3307)/goalsApp?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("failed to connect database")
	}

	db.AutoMigrate(&Goal{})

	return DB{db: db}
}

func (ctx DB) Seed() {
	db := ctx.db

	goals := []Goal{
		{
			Name:        "Master Goal",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et faucibus massa. Sed fringilla rutrum nibh quis interdum. Ut rhoncus erat nec interdum ultricies. Quisque facilisis sapien eu ligula pellentesque dignissim. Ut finibus rutrum lectus,",
		},
		{
			Name:        "Sub Goal 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et faucibus massa. Sed fringilla rutrum nibh quis interdum. Ut rhoncus erat nec interdum ultricies. Quisque facilisis sapien eu ligula pellentesque dignissim. Ut finibus rutrum lectus,",
			ParentID:    1,
		},
		{
			Name:        "Sub Goal 2",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et faucibus massa. Sed fringilla rutrum nibh quis interdum. Ut rhoncus erat nec interdum ultricies. Quisque facilisis sapien eu ligula pellentesque dignissim. Ut finibus rutrum lectus,",
			ParentID:    1,
		},
		{
			Name:        "Sub Goal 3",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et faucibus massa. Sed fringilla rutrum nibh quis interdum. Ut rhoncus erat nec interdum ultricies. Quisque facilisis sapien eu ligula pellentesque dignissim. Ut finibus rutrum lectus,",
			ParentID:    1,
		},
		{
			Name:        "Sub Sub Goal 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et faucibus massa. Sed fringilla rutrum nibh quis interdum. Ut rhoncus erat nec interdum ultricies. Quisque facilisis sapien eu ligula pellentesque dignissim. Ut finibus rutrum lectus,",
			ParentID:    3,
		},
		{
			Name:        "Sub Sub Goal 2",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et faucibus massa. Sed fringilla rutrum nibh quis interdum. Ut rhoncus erat nec interdum ultricies. Quisque facilisis sapien eu ligula pellentesque dignissim. Ut finibus rutrum lectus,",
			ParentID:    3,
		},
	}
	db.Create(&goals)
}

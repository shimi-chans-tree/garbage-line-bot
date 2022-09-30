package repositories

import (
	"chofu-line-bot/app/models"
	"fmt"

	"gorm.io/gorm"
)

type GarbageDayRepository interface {
	GetGarbage(garbege *models.GarbageDay, y int, m int, d int) error
}

type garbageDayRepository struct {
	db *gorm.DB
}

func NewGarbageDayRepository(db *gorm.DB) GarbageDayRepository {
	return &garbageDayRepository{db}
}

func (gr *garbageDayRepository) GetGarbage(garbageDays *models.GarbageDay, y int, m int, d int) error {

	if err := gr.db.Where("year = ? AND month = ? AND day = ?", y, m, d).Find(&garbageDays).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

package main

import (
	"chofu-line-bot/app/db"
	"chofu-line-bot/app/models"

	"gorm.io/gorm"
)

func migrate(dbCon *gorm.DB) {
	// Migration実行
	err := dbCon.AutoMigrate(&models.GarbageDay{})
	if err != nil {
		return
	}
}

func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	// migration実行
	migrate(dbCon)
}

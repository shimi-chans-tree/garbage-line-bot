package main

import (
	"chofu-line-bot/app/batch"
	"chofu-line-bot/app/controllers"
	"chofu-line-bot/app/db"
	"chofu-line-bot/app/logic"
	"chofu-line-bot/app/repositories"
	"chofu-line-bot/app/router"
	"chofu-line-bot/app/services"
)

func main() {

	// DB初期化
	db := db.Init()

	// Repository
	garbageDayRepo := repositories.NewGarbageDayRepository(db)
	// Logic
	garbageDayLogic := logic.NewGarbageDayLogic()
	// Service
	garbageDayService := services.NewGarbageDayService(garbageDayRepo, garbageDayLogic)
	// Controller
	appController := controllers.NewAppController()
	garbageDayController := controllers.NewGarbageDayController(garbageDayService)
	// Router
	appRouter := router.NewAppRouter(appController)
	garbageRouter := router.NewGarbageRouter(garbageDayController)
	mainRouter := router.NewMainRouter(appRouter, garbageRouter)
	// WebServer
	mainRouter.StartWebServer()

	// バッチ処理
	batch.SendMessage()
}

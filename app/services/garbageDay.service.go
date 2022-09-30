package services

import (
	"chofu-line-bot/app/logic"
	"chofu-line-bot/app/models"
	"chofu-line-bot/app/repositories"
	"time"
)

type GarbageDayService interface {
	GetGarbage() (models.BaseGarbageDayResponse, error)
}

type garbageDayService struct {
	gr repositories.GarbageDayRepository
	gl logic.GarbageDayLogic
}

func NewGarbageDayService(gr repositories.GarbageDayRepository, gl logic.GarbageDayLogic) GarbageDayService {
	return &garbageDayService{gr, gl}
}

/*
IDに紐づくTodoを取得
*/
func (gs *garbageDayService) GetGarbage() (models.BaseGarbageDayResponse, error) {
	garbageDay := models.GarbageDay{}
	t := time.Now().AddDate(0, 0, 1)
	y := t.Year()
	m := int(t.Month())
	d := t.Day()

	if err := gs.gr.GetGarbage(&garbageDay, y, m, d); err != nil {
		return models.BaseGarbageDayResponse{}, err
	}
	// レスポンス用の構造体に変換
	responseGarbageDay := gs.gl.CreateGabageDayResponse(&garbageDay)

	return responseGarbageDay, nil
}

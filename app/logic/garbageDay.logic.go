package logic

import "chofu-line-bot/app/models"

type GarbageDayLogic interface {
	CreateGabageDayResponse(garbageDay *models.GarbageDay) models.BaseGarbageDayResponse
}

type garbageDayLogic struct{}

func NewGarbageDayLogic() GarbageDayLogic {
	return &garbageDayLogic{}
}

func (gl *garbageDayLogic) CreateGabageDayResponse(garbageDay *models.GarbageDay) models.BaseGarbageDayResponse {
	var responseGarbageDay models.BaseGarbageDayResponse
	responseGarbageDay.Garbage = garbageDay.Garbage

	return responseGarbageDay
}

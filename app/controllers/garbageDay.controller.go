package controllers

import (
	"chofu-line-bot/app/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type GarbageDayController interface {
	Callback(w http.ResponseWriter, r *http.Request)
}

type garbageDayController struct {
	gs services.GarbageDayService
}

func NewGarbageDayController(gs services.GarbageDayService) GarbageDayController {
	return &garbageDayController{gs}
}

func (gc *garbageDayController) Callback(w http.ResponseWriter, req *http.Request) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := gc.gs.GetGarbage()

	if err != nil {
		return
	}

	var response string

	if res.Garbage != "" {
		response = "☆" + res.Garbage
	} else {
		response = "明日はゴミの日ではありません"
	}

	replyText := "明日"
	errMessage := "「明日」と入力してください"

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyMessage := message.Text
				if strings.Contains(replyMessage, replyText) {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do(); err != nil {
						log.Print(err)
					}
				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			default:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

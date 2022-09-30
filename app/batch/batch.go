package batch

import (
	"chofu-line-bot/app/db"
	"chofu-line-bot/app/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/robfig/cron"
)

func sendMessage() {
	db := db.Init()
	garbageDays := models.GarbageDay{}
	t := time.Now().AddDate(0, 0, 1)
	y := t.Year()
	m := int(t.Month())
	d := t.Day()

	db.Where("year = ? AND month = ? AND day = ?", y, m, d).Find(&garbageDays)

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	var ms string
	
	if garbageDays.Garbage != "" {
		ms = "☆" + garbageDays.Garbage
	} else {
		ms = "明日はゴミの日ではありません"
	}

	message := linebot.NewTextMessage(ms)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

func SendMessage() {
	c := cron.New()
	c.AddFunc("@every 1m", func() { sendMessage() })
	c.Start()
}

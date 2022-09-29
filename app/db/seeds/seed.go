package main

import (
	"chofu-line-bot/app/db"
	"chofu-line-bot/app/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

func garbageSeeds(db *gorm.DB) error {

	webPage := ("https://www.city.chofu.tokyo.jp/www/contents/1646875289947/simple/text20221.txt")
	resp, err := http.Get(webPage)
	if err != nil {
		log.Printf("failed to get html: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("failed to load html: %s", err)
	}

	body := doc.Find("body").Text()
	year := strings.Split(body, "令和")[2:]

	for i := 0; i < len(year); i++ {
		yearGarbage := year[i]

		t := strings.TrimSpace(yearGarbage)
		days := strings.Split(t, "\n")

		for j := 1; j < len(days); j++ {
			month, _ := strconv.Atoi(strings.TrimSpace(days[0][7:8]))
			fmt.Println(month)
			dayArray := strings.Split(days[j], "日")
			day, _ := strconv.Atoi(dayArray[0])
			dayOfWeek := strings.Replace(dayArray[1], "曜", "", -1)
			garbage := dayArray[2]
			garbageDays := models.GarbageDay{
				Year:      2022,
				Month:     month,
				Day:       day,
				DayOfWeek: dayOfWeek,
				Garbage:   garbage,
			}
			if err := db.Create(&garbageDays).Error; err != nil {
				fmt.Printf("%+v", err)
			}
		}
	}
	return nil
}

func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	if err := garbageSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}

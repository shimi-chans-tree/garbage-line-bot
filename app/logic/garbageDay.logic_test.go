package logic_test

import (
	"chofu-line-bot/app/logic"
	"chofu-line-bot/app/models"
	"reflect"
	"testing"
)

func TestCreateGarbageDayResponseSuccess(t *testing.T) {
	// テスト対象の引数
	var argGarbageDay models.GarbageDay
	argGarbageDay.Garbage = "テスト1"

	// 予測値
	var expectedBaseGarbageDayResponse models.BaseGarbageDayResponse
	expectedBaseGarbageDayResponse.Garbage = "テスト1"

	tr := logic.NewGarbageDayLogic()
	// テスト対象の処理を実行
	actual := tr.CreateGabageDayResponse(&argGarbageDay)

	// テスト実行
	if !reflect.DeepEqual(actual, expectedBaseGarbageDayResponse) {
		t.Errorf("actual %v\nwant %v", actual, expectedBaseGarbageDayResponse)
	}
}

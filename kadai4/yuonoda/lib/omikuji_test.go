package omikuji_test

import (
	"fmt"
	omiikuji "github.com/yuonoda/gopherdojo-studyroom/kadai4/yuonoda/lib"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	// テストリクエストを送信
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omiikuji.ExpHandler(w, r)

	// レスポンスを判定
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
}

func TestPickResult(t *testing.T) {
	cases := []struct {
		name          string
		time          time.Time
		expectedRegex string
	}{
		{
			"new_year",
			time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			`大吉`,
		},
		{
			"normal_day",
			time.Now(),
			`^凶$|^吉$|^中吉$|^大吉$`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := omiikuji.ExpPickResult(c.time)
			matched, err := regexp.MatchString(c.expectedRegex, res)
			if err != nil {
				t.Error(err)
			} else if !matched {
				t.Error(fmt.Errorf("result mismatch expected %s, got %s", c.expectedRegex, res))
			}

		})
	}
}

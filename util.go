package telphin

import "time"

//
const (
	API_URL        = "https://apiproxy.telphin.ru"
	API_VERSION    = "/api/ver1.0"
	DATE_FORMAT    = "2006-01-02 15:04:05"
	EXPIRES_OFFSET = 400
	API            = API_URL + API_VERSION
)

//
var hangup = map[string]string{
	"ANSWER":      "вызов был отвечен",
	"BUSY":        "вызов получил сигнал - занято",
	"NOANSWER":    "вонок не отвечен (истек таймер ожидания на сервере)",
	"CANCEL":      "звонящий отменил вызов до истечения таймера ожидания на сервере",
	"CONGESTION":  "произошла ошибка во время вызова",
	"CHANUNAVAIL": "у вызываемого абонента отсутствует регистрация",
}

func HeartBeat(refresh func() (*OAuth, error)) {
	interval := oauth.ExpiresIn - EXPIRES_OFFSET
	for range time.Tick(time.Second * time.Duration(interval)) {
		refresh()
		break
	}
}

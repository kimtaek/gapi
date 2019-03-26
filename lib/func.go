package lib

import (
	"github.com/uniplaces/carbon"
	"log"
	"time"
)

func GetDates(sDate string, eDate string, format string) (data []string) {
	if format == "" {
		format = "2006-01-02T15:04:05Z07:00"
	}
	s, err := carbon.Parse(time.RFC3339, sDate, "Asia/Shanghai")
	e, err := carbon.Parse(time.RFC3339, eDate, "Asia/Shanghai")
	diffDays := s.DiffInDays(e.SubDay(), true)
	var i int64
	var date string
	for i = 0; i <= diffDays; i++ {
		date = s.AddDays(int(i)).Format(format)
		data = append(data, date)
		if date == e.Format(format) {
			break
		}
	}
	if err != nil {
		log.Println(err.Error())
	}
	return data
}

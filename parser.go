package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	WeekNumber   int       `json:"week_number"`
	UtcOffset    string    `json:"utc_offset"`
	UtcDatetime  time.Time `json:"utc_datetime"`
	Unixtime     int       `json:"unixtime"`
	Timezone     string    `json:"timezone"`
	RawOffset    int       `json:"raw_offset"`
	DstUntil     time.Time `json:"dst_until"`
	DstOffset    int       `json:"dst_offset"`
	DstFrom      time.Time `json:"dst_from"`
	Dst          bool      `json:"dst"`
	DayOfYear    int       `json:"day_of_year"`
	DayOfWeek    int       `json:"day_of_week"`
	Datetime     time.Time `json:"datetime"`
	Abbreviation string    `json:"abbreviation"`
}

// get time either from internet or computer depending if connection is available
func getCurrentTime() time.Time {
	if !isConnected() {
		return convertToClock(time.Now())
	}

	// get real current time (in case changes to computer clock)
	resp, err := http.Get(worldClockAPI)
	if err != nil {
		log.Fatal(err)
	}

	tr := TimeResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &tr); err != nil {
		log.Fatal(err)
	}
	return convertToClock(tr.Datetime)
}

func convertToClock(input time.Time) time.Time {
	timeNow := input.Format(timeFormat)
	t, err := time.Parse(timeFormat, timeNow)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

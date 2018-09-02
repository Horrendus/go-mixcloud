package go_mixcloud

import (
	"time"
)

type Feed struct {
	Data []FeedData
}

type FeedData struct {
	URL       	string		`json:"url"`
	CreatedTime	time.Time	`json:"created_time,string"`
	Key			string 		`json:"key"`
}

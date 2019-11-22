package utils

import (
	"strconv"
	"time"
)

func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}

func SendDeleteVideoRequest(id string) {
	//addr := config.GetLbAddr() + ":9001"
	//url := "http://" + addr + "/video-delete-record/" + id
	//_, err := http.Get(url)
	//if err != nil {
	//	log.Printf("Sending deleting video request error: %s", err)
	//}
}

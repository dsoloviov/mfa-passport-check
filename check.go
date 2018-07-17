package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const URL = "http://passport.mfa.gov.ua/"

type Status struct {
	UserSessionId int
	StatusInfo    []HistoryEntry
}

type HistoryEntry struct {
	StatusName   string
	StatusDateUF int
}

func main() {
	id := os.Args[1]

	var body = request(id)

	var status Status

	json.Unmarshal([]byte(body), &status)
	fmt.Printf("Data for SESSION ID %d\n", status.UserSessionId)

	for _, elem := range status.StatusInfo {
		t := formatDate(elem.StatusDateUF)
		fmt.Printf("%s %s\n", t, elem.StatusName)
	}
}

func request(id string) []byte {
	res, err := http.Get(fmt.Sprintf("%s/Home/CurrentSessionStatus?sessionId=%s", URL, id))
	if err != nil {
		log.Fatal(err)
	}
	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err)
	}

	return body
}

func formatDate(timestamp int) time.Time {
	i := int64(timestamp / 1000)

	return time.Unix(i, 0)
}

package go_mixcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestParsingJsonInFeedStruct(t *testing.T) {
	file, e := ioutil.ReadFile("./testdata/feed.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var feed Feed
	err := json.NewDecoder(strings.NewReader(string(file))).Decode(&feed)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for _, feedData := range feed.Data {
			fmt.Println(feedData.Key, feedData.URL, feedData.CreatedTime)
		}
	}
}

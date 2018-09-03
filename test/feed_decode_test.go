/*
 *    Copyright (C) 2018 Stefan Derkits <stefan@derkits.at>
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/horrendus/go-mixcloud/mixcloud"
)

func TestParsingFeedStruct(t *testing.T) {
	var feed mixcloud.Feed
	readFileAndDecodeToFeed("./testdata/feed.json", &feed, t)
	fmt.Println("Paging Prev:", feed.Paging.PreviousURL)
	fmt.Println("Paging Next:", feed.Paging.NextURL)
	fmt.Println("Data Length:", len(feed.Data))
}

func TestParsingFeedStructWithoutNext(t *testing.T) {
	var feed mixcloud.Feed
	readFileAndDecodeToFeed("./testdata/feed_without_next.json", &feed, t)
	fmt.Println("Paging Prev:", feed.Paging.PreviousURL)
	fmt.Println("Paging Next:", feed.Paging.NextURL, feed.Paging.NextURL == "")
	fmt.Println("Data Length:", len(feed.Data))
}

func TestParsingFeedStructEmpty(t *testing.T) {
	var feed mixcloud.Feed
	readFileAndDecodeToFeed("./testdata/feed_empty.json", &feed, t)
	fmt.Println("Paging Prev:", feed.Paging.PreviousURL, feed.Paging.PreviousURL == "")
	fmt.Println("Paging Next:", feed.Paging.NextURL, feed.Paging.NextURL == "")
	fmt.Println("Data Length:", len(feed.Data))
}

func readFileAndDecodeToFeed(fileName string, feed *mixcloud.Feed, t *testing.T) {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	err := json.NewDecoder(strings.NewReader(string(file))).Decode(feed)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
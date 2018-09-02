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

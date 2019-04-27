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

package mixcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestParsingCloudcastsStruct(t *testing.T) {
	var Cloudcasts Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL)
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL)
	fmt.Println("Data Length:", len(Cloudcasts.Data))
	fmt.Printf("%+v\n", Cloudcasts.Data[0])
}

func TestParsingCloudcastsStructWithoutNext(t *testing.T) {
	var Cloudcasts Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts_without_next.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL)
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL, "Next Empty:", Cloudcasts.Paging.NextURL == "")
	fmt.Println("Data Length:", len(Cloudcasts.Data))
}

func TestParsingCloudcastsStructEmpty(t *testing.T) {
	var Cloudcasts Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts_empty.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL, Cloudcasts.Paging.PreviousURL == "")
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL, Cloudcasts.Paging.NextURL == "")
	fmt.Println("Data Length:", len(Cloudcasts.Data))
}

func TestParsingCloudcastsStructEmpty2(t *testing.T) {
	var Cloudcasts Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts_empty_2.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL, Cloudcasts.Paging.PreviousURL == "")
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL, Cloudcasts.Paging.NextURL == "")
	fmt.Println("Data Length:", len(Cloudcasts.Data))
}

func TestRetrieveCloudcasts(t *testing.T) {
	c := NewClient(nil)
	var opt ListOptions
	opt.Since = time.Unix(1476230400, 0)
	opt.Until = time.Unix(1496230400, 0)
	cloudcasts1, err := c.GetCloudcasts("zanjradio", &opt)
	fmt.Println("Error:", err)
	fmt.Println(cloudcasts1)
	fmt.Println("FeedData Length:", len(cloudcasts1.Data))
}

func TestRetrieveAllCloudcasts(t *testing.T) {
	c := NewClient(nil)
	var allCloudcastsData []CloudcastData
	pages := 1
	cloudcasts, err := c.GetCloudcasts("zanjradio", nil)
	if err != nil {
		t.Fail()
	}
	allCloudcastsData = append(allCloudcastsData, cloudcasts.Data...)
	nextUrl := cloudcasts.Paging.NextURL
	for {
		// the following line is necessary to always create a new object, else just some values would be overwritten
		// and missing values would stay the same as before
		cloudcasts = Cloudcasts{}
		err := c.GetPage(nextUrl, &cloudcasts)
		if err != nil {
			t.Fail()
		}
		pages++
		nextUrl = cloudcasts.Paging.NextURL
		if nextUrl == "" {
			break
		}
		allCloudcastsData = append(allCloudcastsData, cloudcasts.Data...)
	}
	fmt.Println("Total Pages", pages)
	fmt.Println("Total Length:", len(allCloudcastsData))
}

func readFileAndDecodeToCloudcasts(fileName string, Cloudcasts *Cloudcasts, t *testing.T) {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	err := json.NewDecoder(strings.NewReader(string(file))).Decode(Cloudcasts)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

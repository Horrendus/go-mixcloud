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

func TestParsingCloudcastsStruct(t *testing.T) {
	var Cloudcasts mixcloud.Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL)
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL)
	fmt.Println("Data Length:", len(Cloudcasts.Data))
}

func TestParsingCloudcastsStructWithoutNext(t *testing.T) {
	var Cloudcasts mixcloud.Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts_without_next.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL)
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL, Cloudcasts.Paging.NextURL == "")
	fmt.Println("Data Length:", len(Cloudcasts.Data))
}

func TestParsingCloudcastsStructEmpty(t *testing.T) {
	var Cloudcasts mixcloud.Cloudcasts
	readFileAndDecodeToCloudcasts("./testdata/cloudcasts_empty.json", &Cloudcasts, t)
	fmt.Println("Paging Prev:", Cloudcasts.Paging.PreviousURL, Cloudcasts.Paging.PreviousURL == "")
	fmt.Println("Paging Next:", Cloudcasts.Paging.NextURL, Cloudcasts.Paging.NextURL == "")
	fmt.Println("Data Length:", len(Cloudcasts.Data))
}

func readFileAndDecodeToCloudcasts(fileName string, Cloudcasts *mixcloud.Cloudcasts, t *testing.T) {
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
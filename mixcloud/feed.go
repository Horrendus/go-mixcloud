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
	"fmt"
	"time"
)

func (c* Client) GetFeed(name string) (*Feed, error) {
	url := fmt.Sprintf("%v/feed/", name)
	fmt.Println(url)
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var feed Feed
	_, err = c.do(req, &feed)
	return &feed, err
}

type Feed struct {
	Paging	Link		`json:"paging,Link"`
	Data	[]FeedData
}

type FeedData struct {
	URL       	string		`json:"url"`
	CreatedTime	time.Time	`json:"created_time,string"`
	Key			string 		`json:"key"`
}

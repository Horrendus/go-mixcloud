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
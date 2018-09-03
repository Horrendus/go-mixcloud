package test

import (
	"fmt"
	"testing"

	"github.com/horrendus/go-mixcloud/mixcloud"
)

func TestRetrieveFeed(t *testing.T) {
	c := mixcloud.NewClient(nil)
	feed1, err := c.GetFeed("zanjradio")
	fmt.Println("Error:", err)
	fmt.Println(feed1)
	fmt.Println("FeedData Length:", len(feed1.Data))
}

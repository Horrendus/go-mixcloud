package test

import (
	"fmt"
	"testing"

	"github.com/horrendus/go-mixcloud/mixcloud"
)

func TestRetrieveCloudcasts(t *testing.T) {
	c := mixcloud.NewClient(nil)
	cloudcasts1, err := c.GetCloudcasts("zanjradio")
	fmt.Println("Error:", err)
	fmt.Println(cloudcasts1)
	fmt.Println("FeedData Length:", len(cloudcasts1.Data))
}

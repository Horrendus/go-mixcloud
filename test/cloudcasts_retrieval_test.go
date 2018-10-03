package test

import (
	"fmt"
	"testing"

	"github.com/horrendus/go-mixcloud/mixcloud"
)

func TestRetrieveCloudcasts(t *testing.T) {
	c := mixcloud.NewClient(nil)
	cloudcasts1, err := c.GetCloudcasts("zanjradio", nil)
	fmt.Println("Error:", err)
	fmt.Println(cloudcasts1)
	fmt.Println("FeedData Length:", len(cloudcasts1.Data))
}

func TestRetrieveAllCloudcasts(t *testing.T) {
	c := mixcloud.NewClient(nil)
	var allCloudcastsData []mixcloud.CloudcastData
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
		cloudcasts = mixcloud.Cloudcasts{}
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

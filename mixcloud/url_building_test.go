package mixcloud

import (
	"fmt"
	"testing"
	"time"
)

func TestBuildListOptions(t *testing.T) {
	var opt ListOptions
	opt.Offset = 100
	opts, err := addListOptions("http://test.at", &opt)
	if err != nil {
		t.Fail()
	}
	fmt.Println(opts)
}

func TestBuildListOptionsWithLimit(t *testing.T) {
	var opt ListOptions
	opt.Limit = 30
	opts, err := addListOptions("http://test.at", &opt)
	if err != nil {
		t.Fail()
	}
	fmt.Println(opts)
}

func TestBuildListOptionsIgnoredOffset(t *testing.T) {
	var opt ListOptions
	opt.Offset = 100
	opt.Since = time.Unix(1476230400,0)
	opt.Until = time.Unix(1496230400,0)
	opts, err := addListOptions("http://test.at", &opt)
	if err != nil {
		t.Fail()
	}
	fmt.Println(opts)
}

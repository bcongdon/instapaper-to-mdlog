package main

import (
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
)

var (
	fakeItem1 = &gofeed.Item{Title: "Title1", Link: "http://link1"}
	fakeItem2 = &gofeed.Item{Title: "Title2", Link: "http://link2"}
	fakeItem3 = &gofeed.Item{Title: "Title3", Link: "http://link3?q=123"}
)

func TestMergeItems(t *testing.T) {
	date := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local)
	type args struct {
		orig  string
		date  time.Time
		items []*gofeed.Item
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty orig",
			args{"", date, []*gofeed.Item{fakeItem1, fakeItem2}},
			"# January 1, 2021\n" +
				"- [Title1](http://link1)\n" +
				"- [Title2](http://link2)\n\n",
		},
		{
			"orig with previous date",
			args{
				"# December 31, 2020\n- [Foo](http://bar)",
				date,
				[]*gofeed.Item{fakeItem1},
			},
			"# January 1, 2021\n" +
				"- [Title1](http://link1)\n\n" +
				"# December 31, 2020\n- [Foo](http://bar)",
		},
		{
			"orig with existing links for current date",
			args{
				"# January 1, 2021\n- [Foo](http://bar)",
				date,
				[]*gofeed.Item{fakeItem1},
			},
			"# January 1, 2021\n" +
				"- [Title1](http://link1)\n" +
				"- [Foo](http://bar)",
		},
		{
			"link with meta characters",
			args{
				"# January 1, 2021\n- [Foo](http://link3?q=123)",
				date,
				[]*gofeed.Item{fakeItem3},
			},
			"# January 1, 2021\n" +
				"- [Foo](http://link3?q=123)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeItems(tt.args.orig, tt.args.date, tt.args.items)
			if err != nil {
				t.Fatalf("mergeItems() got error: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("mergeItems() = %q, want %q", got, tt.want)
			}
		})
	}
}

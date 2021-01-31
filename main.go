package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

var (
	feedURL = flag.String("feedURL", "", "The Instapaper feed to save")
	outFile = flag.String("out", "", "The file to save the feed")
)

func main() {
	flag.Parse()
	if *feedURL == "" || *outFile == "" {
		log.Fatal("-feedURL and -out are required")
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(*feedURL)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadFile(*outFile)
	merged, err := mergeItems(string(contents), time.Now(), feed.Items)

	ioutil.WriteFile(*outFile, []byte(merged), os.ModePerm)
}

func mergeItems(orig string, date time.Time, items []*gofeed.Item) (string, error) {
	var unsavedItems []*gofeed.Item
	for _, item := range items {
		itemRegex, err := regexp.Compile(fmt.Sprintf(`\[.*\]\(%s\)`, regexp.QuoteMeta(item.Link)))
		if err != nil {
			return "", nil
		}
		if !itemRegex.MatchString(orig) {
			unsavedItems = append(unsavedItems, item)
		}
	}
	// If nothing to add, bail early.
	if len(unsavedItems) == 0 {
		return orig, nil
	}

	var itemsStrBuilder strings.Builder
	for idx, item := range unsavedItems {
		itemsStrBuilder.WriteString(fmt.Sprintf("- [%s](%s)", item.Title, item.Link))
		if idx != len(unsavedItems)-1 {
			itemsStrBuilder.WriteByte('\n')
		}
	}
	itemsStr := itemsStrBuilder.String()

	header := dateHeader(date)
	// If we're adding a "new day", prepend it to the existing items.
	if !strings.Contains(orig, header) {
		return fmt.Sprintf("%s\n%s\n\n%s", header, itemsStr, orig), nil
	}

	// If we're modifying an original day, add it immediately after the date header.
	idx := strings.Index(orig, header) + len(header)
	return orig[:idx] + "\n" + itemsStr + orig[idx:], nil
}

func dateHeader(date time.Time) string {
	return date.Format("# January 2, 2006")
}

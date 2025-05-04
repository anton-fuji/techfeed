package feed

import (
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v2"
)

type FeedItem struct {
	Title     string
	Link      string
	Published time.Time
	FeedLink  string
	Source    string
}

type FeedGroup struct {
	URLs     []string `yaml:"urls"`
	Template string   `yaml:"template"`
	Limit    int      `yaml:"limit"`
	Sorted   bool     `yaml:"sorted"`
}

func LoadFeeds(path string) ([]FeedGroup, map[string][]FeedItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	var groups []FeedGroup
	if err := yaml.Unmarshal(data, &groups); err != nil {
		return nil, nil, err
	}

	fp := gofeed.NewParser()
	groupedItems := make(map[string][]FeedItem)

	for _, group := range groups {
		var items []FeedItem

		for _, url := range group.URLs {
			feed, err := fp.ParseURL(url)
			if err != nil {
				log.Printf("Feedの取得に失敗(%s): %v", url, err)
				continue
			}

			source := "unknown"
			switch {
			case strings.Contains(feed.FeedLink, "qiita.com"):
				source = "qiita"
			case strings.Contains(feed.FeedLink, "zenn.dev"):
				source = "zenn"
			}

			for _, entry := range feed.Items {
				if entry.PublishedParsed == nil {
					continue
				}
				items = append(items, FeedItem{
					Title:     entry.Title,
					Link:      entry.Link,
					Published: *entry.PublishedParsed,
					FeedLink:  feed.FeedLink,
					Source:    source,
				})
			}
		}

		if group.Sorted {
			sort.Slice(items, func(i, j int) bool {
				return items[i].Published.After(items[j].Published)
			})
		}

		if group.Limit > 0 && len(items) > group.Limit {
			items = items[:group.Limit]
		}

		groupedItems[group.Template] = items
	}

	return groups, groupedItems, nil
}

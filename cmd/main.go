package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/anton-fuji/techfeed/feed"
)

func main() {
	groups, groupedItems, err := feed.LoadFeeds("config/feeds-setup.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range groups {
		items := groupedItems[group.Template]

		tmpl, err := template.ParseFiles(group.Template)
		if err != nil {
			log.Printf("テンプレートが読み込めません: %s (%v)", group.Template, err)
			continue
		}

		fmt.Printf("===== %s =====\n", group.Template)
		err = tmpl.Execute(os.Stdout, items)
		if err != nil {
			log.Printf("テンプレートの実行に失敗: %v", err)
		}

		fmt.Println()
	}
}

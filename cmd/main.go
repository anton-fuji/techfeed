package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"

	"github.com/anton-fuji/techfeed/feed"
)

func main() {
	_, groupedItems, err := feed.LoadFeeds("config/feeds-setup.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var renderOutput bytes.Buffer

	for _, items := range groupedItems {
		tmpl, err := template.ParseFiles("templates/feeds.tmpl")
		if err != nil {
			log.Printf("テンプレートが読み込めません: %v", err)
			continue
		}

		if err := tmpl.Execute(&renderOutput, items); err != nil {
			log.Printf("テンプレートの実行に失敗: %v", err)
		}
		renderOutput.WriteString("\n")
	}

	readmeBytes, err := os.ReadFile("README.md")
	if err != nil {
		log.Fatalf("README.mdの読み込みに失敗: %v", err)
	}
	readme := string(readmeBytes)

	re := regexp.MustCompile(`(?s)<!-- feeds:start -->.*?<!-- feeds:end -->`)
	newContent := fmt.Sprintf("<!-- feeds:start -->\n%s<!-- feeds:end -->", renderOutput.String())
	updated := re.ReplaceAllString(readme, newContent)

	if err := os.WriteFile("README.md", []byte(updated), 0664); err != nil {
		log.Fatalf("README.md の書き込みに失敗: %v", err)
	}
	fmt.Println("✔ README.mdを更新しました!")
}

### ディレクトリ構成

```golang
techfeed/
├── cmd/
│   └──  main.go 　           # エントリーポイント（main関数）
├── config/
│   └── feeds.yaml             # RSSフィードの設定（YAML形式）
├── feed/
│   ├── parser.go              # RSSを取得・パースするロジック
│   └── model.go               # FeedConfigなど構造体の定義
├── output/
│   └── formatter.go           # Markdownなどの整形出力用処理
├── go.mod
└── go.sum
```
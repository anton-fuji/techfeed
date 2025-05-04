package feed

type Feeds struct {
	Urls     []string `yaml:"urls"`
	Template string   `yaml:"template"`
	Limit    int      `yaml:"limit"`
	Sorted   bool     `yaml:"sorted"`
}

// type Config struct {
// 	Feeds []FeedConfig `yaml:"feeds"`
// }

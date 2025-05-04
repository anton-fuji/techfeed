package feed

type Feeds struct {
	Title string `yaml:""title`
	Url   string `yaml: "url"`
	Max   string `yaml: "url"`
}

type Config struct {
	Feeds []FeedConfig `yaml:"feeds"`
}

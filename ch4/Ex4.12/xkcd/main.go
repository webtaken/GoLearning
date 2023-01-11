package xkcd

const ComicURL = "https://xkcd.com"

type Comic struct {
	Title      string `json:"title"`
	Link       string `json:"link"`
	Transcript string `json:"transcript"`
	Img        string `json:"img"`
}

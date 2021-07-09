package src

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/riskiramdan/tiktok-web-scrapper/models"
)

type Tiktok struct {
	scraper *colly.Collector
}

func (t *Tiktok) GetVideoProperties(copiedLink string) (data *models.TiktokResponseData, err error) {
	if isTiktokUrl(copiedLink) {
		dataInterface := models.Data{}
		var err error
		t.scraper.OnHTML("script[id=__NEXT_DATA__]", func(e *colly.HTMLElement) {
			err = json.Unmarshal([]byte(e.Text), &dataInterface)
		})
		err = t.scraper.Visit(copiedLink)

		data := &models.TiktokResponseData{
			Author:      dataInterface.Props.PageProps.ItemInfo.ItemStruct.Author,
			AuthorStats: dataInterface.Props.PageProps.ItemInfo.ItemStruct.AuthorStats,
			Stats:       dataInterface.Props.PageProps.ItemInfo.ItemStruct.Stats,
			Video:       dataInterface.Props.PageProps.ItemInfo.ItemStruct.Video,
		}
		return data, err
	}
	return nil, errors.New("invalid tiktok url")
}

func NewTiktok() *Tiktok {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Host", "v16-web.tiktokcdn.com")
		r.Headers.Set("Sec-Fetch-Dest", "video")
		r.Headers.Set("Sec-Fetch-Mode", "no-cors")
		r.Headers.Set("Sec-Fetch-Site", "cross-site")
	})
	c.AllowURLRevisit = true
	t := &Tiktok{
		scraper: c,
	}
	return t
}

func isTiktokUrl(str string) bool {
	if strings.Contains(str, "tiktok") {
		u, err := url.Parse(str)
		return err == nil && u.Scheme != "" && u.Host != ""
	}
	return false
}

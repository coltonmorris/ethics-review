package methods

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

import "github.com/gocolly/colly"

func Bing(r []string, q string) []float64 {
	cnts := make(chan *resp)

	for i := 0; i < 3; i++ {
		c := colly.NewCollector(
			colly.AllowURLRevisit(),
			colly.IgnoreRobotsTxt(),
		)
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		n := r[i]
		c.OnError(func(_ *colly.Response, err error) {
			fmt.Println("Something went wrong:", err)
			cnts <- &resp{Name: n, Count: 0}
		})

		c.OnHTML("#b_content .sb_count", func(e *colly.HTMLElement) {
			c := e.Text
			c2 := strings.Replace(c, ",", "", 10)
			c3 := strings.Split(c2, " ")
			c4, _ := strconv.Atoi(c3[0])
			cnts <- &resp{Name: n, Count: c4}
		})

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", RandomString())
		})

		addrPath := "http://www.bing.com/search?q="
		addrQuery := r[i] + " " + q
		go c.Visit(addrPath + url.QueryEscape(addrQuery))
	}

	return getNormalResponses(cnts, r)
}

var startBing StartFunc = func(qna *QandA, doneChannel chan *MethodResults) {
	var bingMethod = &MethodResults{
		Name:    "BingResults",
		Results: Bing(qna.Answers, qna.Question),
	}

	doneChannel <- bingMethod
	fmt.Println("Bing")
}

package methods

import (
  "fmt"
  "strings"
  "strconv"
  "net/url"
)

import "github.com/gocolly/colly"


type resp struct {
  Name string
  Count int
}


func Bing(r []string, q string) []float64 {
  cnts := make(chan *resp)

  for i := 0; i < 3; i++ {
    c := colly.NewCollector()
    c.OnRequest(func(r *colly.Request) {
      fmt.Println("Visiting", r.URL)
    })

    n := r[i]
    c.OnError(func(_ *colly.Response, err error) {
      fmt.Println("Something went wrong:", err)
      cnts <- &resp{Name: n, Count: 0}
    })
    
    c.OnHTML("#b_content", func(e *colly.HTMLElement) {
      c :=  e.ChildText(".sb_count")
      c2 :=strings.Replace(c,",", "", 10)
      c3 := strings.Split(c2, " ")
      c4, _ := strconv.Atoi(c3[0])
      cnts <- &resp{Name: n, Count: c4}
    })

    addrPath := "https://www.bing.com/search?q="
    addrQuery :=  r[i] + " " + q
    go c.Visit(addrPath + url.QueryEscape(addrQuery))
  }
  
  results := []*resp{}

  for i := 0; i < 3; i++ {
    q := <-cnts
    fmt.Println(q)
    results = append(results, q)
  }

  var total float64 = 0
  for _, v := range results {
    total += float64(v.Count)
  }

  normalized_results := []float64{}
  for _, v := range r {
    for _, v2 := range results {
      if v == v2.Name {
        normalized_results = append(normalized_results, float64(v2.Count) / total)
      }
    }
  }

  fmt.Println(normalized_results)
  return normalized_results
}


var startBing StartFunc = func(qna *QandA, doneChannel chan *MethodResults) {
  var bingMethod = &MethodResults{
    Name: "BingResults",
    Results: Bing(qna.Answers, qna.Question),
  }

  doneChannel <- bingMethod
  fmt.Println("Bing")
}

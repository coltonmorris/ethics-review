package methods

import "fmt"


var startBing StartFunc = func(qna *QandA, doneChannel chan *MethodResults) {
  var bingMethod = &MethodResults{
    Name: "BingResults",
    Results: []float64{0.7,0.2,0.1}}

  doneChannel <- bingMethod
  fmt.Println("Bing")
}

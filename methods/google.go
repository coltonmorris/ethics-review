package methods

import "fmt"

var startGoogle StartFunc = func(qna *QandA, doneChannel chan *MethodResults) {
	var googleMethod = &MethodResults{
		Name:    "GoogleResults",
		Results: []float64{0.9, 0.09, 0.01}}

	doneChannel <- googleMethod
	fmt.Println("Google")
}

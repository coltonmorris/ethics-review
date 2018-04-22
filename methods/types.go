package methods

type QandA struct {
	Question string
	Answers  []string
}

type MethodResults struct {
	Name string
	// each index corresponds to the index of the answer
	Results []float64
}

type StartFunc func(qna *QandA, doneChannel chan *MethodResults)

func Start(s StartFunc, qna *QandA, doneChannel chan *MethodResults) { s(qna, doneChannel) }

type QuorumResults struct {
	Qna         *QandA
	Methods     []*MethodResults
	FinalAnswer []float64
}

var Methods []*MethodResults
var StartMethods []StartFunc

type resp struct {
	Name  string
	Count int
}



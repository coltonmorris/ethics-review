package main

import m "github.com/coltonmorris/ethics-review/methods"

type QandA struct {
	question string
	answers  []string
}

type QuorumResults struct {
  qna *QandA
  methods []*m.MethodResults
  finalAnswer []float64
}



func runQuorum(qna *QandA) *QuorumResults {

  quorum := &QuorumResults{
    qna: qna,
    methods: m.Methods,
    // TODO calculate this by doing an average
    finalAnswer: []float64{0.8, 0.15, 0.05}}


  // TODO start here when you finish a method   
  // for _, method := range quorum.methods {
    // go method.start(qna, doneChannel)
  // }
  // listen for a method to finish and print it's results
  // go func() {
  //  doneChannel <- result
  //  PrintQuorumResults(calculateQuorumResults())
  //  
  // }



  return quorum
}


func main() {
  // TODO parse ocr_output.txt to get these
	qna := &QandA{
		question: "Why are you who you are?",
		answers:  []string{"because", "same", "raised by wolves"}}

  quorumResults := runQuorum(qna)

  PrintQuorumResults(quorumResults)
}

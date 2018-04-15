package main

import (
	"fmt"
	m "github.com/coltonmorris/ethics-review/methods"
	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var black = color.New(color.FgBlack).Add(color.Underline).SprintFunc()

func PrintQuorumResults(quorum *m.QuorumResults) {
	PrintQandA(quorum.Qna)
	for _, method := range quorum.Methods {
		PrintMethodResults(quorum.Qna, method)
	}
	PrintFinalAnswer(quorum.Qna, quorum.FinalAnswer)
}

func PrintFinalAnswer(qna *m.QandA, finalAnswer []float64) {
	smallest, middle, largest := IndexOfSmallMiddleLarge(finalAnswer)

	fmt.Println(black("Final Answer"))
	fmt.Printf("\tSmallest: %s\n", red(finalAnswer[smallest], "\t ", qna.Answers[smallest]))
	fmt.Printf("\tMiddle:   %s\n", yellow(finalAnswer[middle], "\t ", qna.Answers[middle]))
	fmt.Printf("\tLargest:  %s\n", green(finalAnswer[largest], "\t ", qna.Answers[largest]))
}

func PrintMethodResults(qna *m.QandA, method *m.MethodResults) {
	smallest, middle, largest := IndexOfSmallMiddleLarge(method.Results)
  fmt.Println("Indexes of (smallest, middle, largest) : (", smallest, ", ", middle, ", ", largest, ")")

	fmt.Printf("Method: %s\n", blue(method.Name))
	fmt.Printf("\tSmallest: %s\n", red(method.Results[smallest], "\t ", qna.Answers[smallest]))
	fmt.Printf("\tMiddle:   %s\n", yellow(method.Results[middle], "\t ", qna.Answers[middle]))
	fmt.Printf("\tLargest:  %s\n", green(method.Results[largest], "\t ", qna.Answers[largest]))
}

func PrintQandA(qna *m.QandA) {
	fmt.Printf("Question: %s\n", blue(qna.Question))
	fmt.Printf("Answers: \t%s,\n \t\t%s,\n \t\t%s,\n", blue(qna.Answers[0]), blue(qna.Answers[1]), blue(qna.Answers[2]))
}

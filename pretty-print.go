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

func PrintQuorumResults(quorum *QuorumResults) {
	PrintQandA(quorum.qna)
	for _, method := range quorum.methods {
		PrintMethodResults(quorum.qna, method)
	}
	PrintFinalAnswer(quorum.qna, quorum.finalAnswer)
}

func PrintFinalAnswer(qna *QandA, finalAnswer []float64) {
	smallest, middle, largest := IndexOfSmallMiddleLarge(finalAnswer)

	fmt.Println(black("Final Answer"))
	fmt.Printf("\tSmallest: %s\n", red(finalAnswer[smallest], "\t ", qna.answers[smallest]))
	fmt.Printf("\tMiddle:   %s\n", yellow(finalAnswer[middle], "\t ", qna.answers[middle]))
	fmt.Printf("\tLargest:  %s\n", green(finalAnswer[largest], "\t ", qna.answers[largest]))
}

func PrintMethodResults(qna *QandA, method *m.MethodResults) {
	smallest, middle, largest := IndexOfSmallMiddleLarge(method.Results)

	fmt.Printf("Method: %s\n", blue(method.Name))
	fmt.Printf("\tSmallest: %s\n", red(method.Results[smallest], "\t ", qna.answers[smallest]))
	fmt.Printf("\tMiddle:   %s\n", yellow(method.Results[middle], "\t ", qna.answers[middle]))
	fmt.Printf("\tLargest:  %s\n", green(method.Results[largest], "\t ", qna.answers[largest]))
}

func PrintQandA(qna *QandA) {
	fmt.Printf("Question: %s\n", blue(qna.question))
	fmt.Printf("Answers: \t%s\n \t\t%s\n \t\t%s\n", blue(qna.answers[0]), blue(qna.answers[1]), blue(qna.answers[2]))
}

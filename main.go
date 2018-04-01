package main

import m "github.com/coltonmorris/ethics-review/methods"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type QandA struct {
	question string
	answers  []string
}

type QuorumResults struct {
	qna         *QandA
	methods     []*m.MethodResults
	finalAnswer []float64
}

func runQuorum(qna *QandA) *QuorumResults {

	quorum := &QuorumResults{
		qna:     qna,
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

func parseQandA(path string) (QandA, error) {
	var qna QandA

	file, err := os.Open(path)
	if err != nil {
		return qna, err
	}
	defer file.Close()

	questionParsed := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Trim(text, " \n\t")

		if len(text) != 0 {
			if !questionParsed {
				qna.question += text
			} else {
				qna.answers = append(qna.answers, text)
			}

			if len(text) != 0 && text[len(text)-1] == '?' {
				questionParsed = true
			}
		}
	}

	return qna, scanner.Err()
}

func main() {
	qna, err := parseQandA("ocr_output.txt")

	if err != nil {
		fmt.Println("ERROR READING OCR OUTPUT:", err)
		return
	}

	quorumResults := runQuorum(&qna)

	PrintQuorumResults(quorumResults)
}

package main

import m "github.com/coltonmorris/ethics-review/methods"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func runQuorum(qna *m.QandA) *m.QuorumResults {
  doneChannel := make(chan *m.MethodResults, len(m.StartMethods))

  // start methods
  for _, method := range m.StartMethods {
    go m.Start(method, qna, doneChannel)
  }

  var methods []*m.MethodResults

  for i := 1; i <= len(m.StartMethods); i++ {
    method := <-doneChannel
    methods = append(methods, method)
  }

  quorum := &m.QuorumResults{
    Qna:     qna,
    Methods: methods,
    // TODO calculate this by doing an average
    FinalAnswer: []float64{0.8, 0.15, 0.05}}

  close(doneChannel)
  return quorum
}

func parseQandA(path string) (m.QandA, error) {
	var qna m.QandA

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
				qna.Question += text
			} else {
				qna.Answers = append(qna.Answers, text)
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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	m "github.com/coltonmorris/ethics-review/methods"
	"github.com/manifoldco/promptui"
)

func calculateFinalAnswer(methods []*m.MethodResults) []float64 {
	finalAnswer := []float64{0, 0, 0}
	total := []float64{0, 0, 0}
	for _, method := range methods {
		for i := 0; i < 3; i++ {
			total[i] += method.Results[i]
		}
	}

	// calculate average
	var sumOfTotal float64 = 0
	for i := 0; i < 3; i++ {
		if total[i] != 0 {
			finalAnswer[i] = total[i] / 3
			sumOfTotal += finalAnswer[i]
		}
	}

	// normalize
	for i := 0; i < len(total); i++ {
		finalAnswer[i] = finalAnswer[i] / sumOfTotal
	}

	return finalAnswer
}

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

	PrintMethodResults(qna, methods[0])

	quorum := &m.QuorumResults{
		Qna:     qna,
		Methods: methods,
		// TODO calculate this by doing an average
		FinalAnswer: calculateFinalAnswer(methods)}

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

		// TODO: handle multiline answers?
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

	PrintQandA(&qna)
	return qna, scanner.Err()
}

func getCorrectAnswer(quorumResults *m.QuorumResults) (int, error) {
	answers := quorumResults.Qna.Answers

	prompt := promptui.Select{
		Label: "Select Correct Answer",
		Items: answers,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}

	for i, ele := range answers {
		if result == ele {
			return i, nil
		}
	}
  fmt.Println("1")
	return -1, errors.New("Index out of range")
}

func main() {
	qna, err := parseQandA("ocr_output.txt")

	if err != nil {
		fmt.Println("ERROR READING OCR OUTPUT:", err)
		return
	}

	quorumResults := runQuorum(&qna)

	PrintQuorumResults(quorumResults)

	correctIndex, err := getCorrectAnswer(quorumResults)
	if err != nil {
    fmt.Println("error getting correct answer: ", correctIndex, err)
    return
	}
  saveQuorumToCsv(quorumResults, correctIndex)
}

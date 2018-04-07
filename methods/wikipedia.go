package methods

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type PageResult struct {
	answer string
	hits   int
}

func GetWikipedia(answer string) (string, error) {
	resp, err := http.Get("http://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=" + answer)
	if err != nil {
		// handle error
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	response := string(body)
	return response, nil
}

func GetNumberOfOccurences(c chan<- PageResult, answer string, nouns []string) {
	count := 0
	wikiBody, _ := GetWikipedia(answer) // TODO: check for error

	if len(wikiBody) < 113*len(answer)*3 { // if its too short, try titlecase
		wikiBody, _ = GetWikipedia(strings.Title(answer))
	}

	for _, noun := range nouns {
		count += strings.Count(wikiBody, noun)
	}

	pr := PageResult{
		answer: answer,
		hits:   count,
	}

	c <- pr
}

func CalculateWeights(results [3]int) []float64 {
	total := 0
	for _, num := range results {
		total += num
	}

	var weights []float64
	for _, val := range results {
		weights = append(weights, float64(val)/float64(total))
	}

	return weights
}

// the function that grabs
var startWikipedia StartFunc = func(qna *QandA, doneChannel chan *MethodResults) {
	var numbers [3]int
	c := make(chan PageResult)

	// TODO: get an actual list of nouns, not just question minus short words
	nouns := strings.Split(qna.Question, " ")
	blen := len(nouns)
	for n := 0; n < blen; n++ {
		if len(nouns[n]) > 3 {
			nouns = append(nouns, nouns[n])
		}
	}
	nouns = nouns[blen:]

	// spin up goroutines to count up
	for i := 0; i < len(qna.Answers); i++ {
		go GetNumberOfOccurences(c, qna.Answers[i], nouns)
	}

	// consume PageResults from channel
	for r := 0; r < 3; r++ {
		result := <-c
		// find the answers array
		for i, val := range qna.Answers {
			if val == result.answer {
				numbers[i] = result.hits
			}
		}
	}

	// calculate the weights
	weights := CalculateWeights(numbers)

	methodResults := &MethodResults{
		Name:    "Wikipedia",
		Results: weights,
	}

	doneChannel <- methodResults
}

package methods

import (
	"fmt"
	"math/rand"
)
func getNormalResponses(c chan *resp, r []string) []float64 {
	results := []*resp{}

  fmt.Println("len of r: ",len(r))
	for i := 0; i < 3; i++ {
		q := <- c
    fmt.Println("recieved result: ",q)
		results = append(results, q)
	}

	var total float64 = 0
	for _, v := range results {
		total += float64(v.Count)
	}

	if total == 0 {
		fmt.Println("no responses")
		return []float64{0,0,0}
	}

	normalized_results := []float64{}
	for _, v := range r {
		for _, v2 := range results {
			if v == v2.Name {
				normalized_results = append(normalized_results, float64(v2.Count)/total)
			}
		}
	}

	fmt.Println(normalized_results)
	return normalized_results
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

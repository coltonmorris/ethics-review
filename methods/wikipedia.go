package methods

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

var wikipediaMethod = &MethodResults{
	Name:    "WikipediaResults",
	Results: []float64{0.2, 0.2, 0.6},
}

// func GetWikipedia(answer string) (string, error) {
// 	resp, err := http.Get("http://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=" + answer)
// 	if err != nil {
// 		// handle error
// 		fmt.Println("Error querying wikipedia", err)
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)

// 	response := string(body)
// 	return response, nil
// }

// if(err != nil){
// 	fmt.Println("Error GetWikipedia returned an error: ", err)
// }

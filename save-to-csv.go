package main

import (
	"bufio"
	"encoding/csv"
  "path/filepath"
	"fmt"
  "log"
  "io"
	"os"
	"strconv"

  m "github.com/coltonmorris/ethics-review/methods"
)

type MethodPrediction struct {
  Name string
  Average float64 
}

type CsvRecord struct {
  question string 
  answer0 string 
  answer1 string 
  answer2 string 
  weight0 float64 
  weight1 float64 
  weight2 float64 
  correctAnswerIndex int 
  guessedAnswerIndex int 
}

func saveQuorumToCsv(quorumResults *m.QuorumResults, correctIndex int) {
	// TODO save the actual quorum prediction
	fmt.Println("Saving results...")

	// save each methods predictions
	for _, method := range quorumResults.Methods {
		filename := "data/" + method.Name + ".csv"
		fields := []string{"question", "answer0", "answer1", "answer2", "weight0", "weight1", "weight2", "correctAnswerIndex", "guessedAnswerIndex"}
		// question and answers should be surrounded in quotes

		// create a csv writer. create the file with headers if it doesn't exist.
    var w *csv.Writer
		if _, err := os.Stat(filename); os.IsNotExist(err) {
      fmt.Println("CSV file did not exist. Creating it...")

      // try to create data directory if it doesn't exist
      newpath := filepath.Join(".", "data")
      os.MkdirAll(newpath, os.ModePerm)

			// touch file
			newFile, err := os.Create(filename)
			defer newFile.Close()
			if err != nil {
				fmt.Println("Error creating new file:", err)
			}

			w = csv.NewWriter(bufio.NewWriter(newFile))

			err = w.Write(fields)
			if err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		} else {
			file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			defer file.Close()
			if err != nil {
				fmt.Println("Error opening file for CSV:", err)
			}
			w = csv.NewWriter(bufio.NewWriter(file))
		}

		record := []string{quorumResults.Qna.Question, quorumResults.Qna.Answers[0], quorumResults.Qna.Answers[1], quorumResults.Qna.Answers[2]}

    var count float64 = 0
		for _, result := range method.Results {
      if result == 0 { count += 1 }
			record = append(record, fmt.Sprintf("%.6f", result))
		}
    
    // ignore records that predict 0 for all the weights
    if count != 3 {
      record = append(record, strconv.Itoa(correctIndex))

      _, _, large := IndexOfSmallMiddleLarge(method.Results)
      record = append(record, strconv.Itoa(large))

      err := w.Write(record)
      if err != nil {
        log.Fatalln("error writing record to csv:", err)
      }

      w.Flush()	// Write any buffered data to the underlying writer (standard output).

      if err := w.Error(); err != nil {
        log.Fatal("yup: ", err)
      }
    }
	}
}

func getMethodPredictionAverages(methods []*m.MethodResults) []*MethodPrediction {
  methodPredictions := []*MethodPrediction{}

  // read from data
  for _, method := range methods {
    filename := "data/" + method.Name + ".csv"
    csvFile, _ := os.Open(filename)
    defer csvFile.Close()
    reader := csv.NewReader(bufio.NewReader(csvFile))

    var lines float64 = 0
    var correctCount float64 = 0
    for {
      line, error := reader.Read()
      if error == io.EOF {
        break
      } else if error != nil {
        log.Fatal(error)
      }

      // weight0, _ := strconv.ParseFloat(line[4], 64)
      // weight1, _ := strconv.ParseFloat(line[5], 64)
      // weight2, _ := strconv.ParseFloat(line[6], 64)
      correctAnswerIndex, _ :=strconv.Atoi(line[7])
      guessedAnswerIndex, _ :=strconv.Atoi(line[8])

      if correctAnswerIndex == guessedAnswerIndex { correctCount += 1 }
      lines += 1
    }

    if lines == 0 { log.Fatal("No data in method") }

    methodPredictions = append(methodPredictions, &MethodPrediction{ Name: method.Name, Average: correctCount/lines })
  }

  return methodPredictions
}

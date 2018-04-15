package main

import "testing"
import (
  "fmt"
  m "github.com/coltonmorris/ethics-review/methods"
)


// mock data
var qna = &m.QandA{
  Question: "What color is NOT in the flag for the united states of america",
  Answers: []string{"red", "white", "yellow"}}
var googleMethod = &m.MethodResults{
  Name:    "GoogleResults",
  // Results: []float64{0.2,0.3,0.5}}
  Results: []float64{0.0,0.0,0.0}}
var bingMethod = &m.MethodResults{
  Name:    "BingResults",
  // Results: []float64{0.1,0.5,0.4}}
  Results: []float64{0.0,0.0,0.0}}
var wikipediaMethod = &m.MethodResults{
  Name:    "Wikipedia",
  Results: []float64{0.05,0.05,0.9}}

var methods []*m.MethodResults = []*m.MethodResults{googleMethod, bingMethod, wikipediaMethod}


func TestCalculateFinalAnswerCsvAverage(t *testing.T) {
  finalAnswer := calculateFinalAnswerCsvAverage(methods)
  fmt.Println(finalAnswer)
}

func TestGetMethodPredictionAverages(t *testing.T) {
  predictions := getMethodPredictionAverages(methods)

  fmt.Println("")
  fmt.Println("Accuracies: ")
  for _, pred := range predictions{
    fmt.Println(pred.Name, ": ", pred.Average*100, "%")
  }
  fmt.Println("")
}

func TestRemoveEmptyMethodResults(t *testing.T) {
  finalMethods := removeEmptyMethodResults(methods)

  for _, method := range finalMethods {
    PrintMethodResults(qna, method)
  }
}

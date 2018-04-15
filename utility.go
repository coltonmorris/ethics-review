package main

import (
  "sort"
)

// finds the smallest, middleIndex, and largestIndex answers
func IndexOfSmallMiddleLarge(arr []float64) (int, int, int) {
  // sort a copy of the array
  sorted := make([]float64, len(arr))
  copy(sorted, arr)
  sort.Float64s(sorted)

  var smallestIndex, middleIndex, largestIndex int = -1,-1,-1
  for i, ele := range arr {
    if ele == sorted[0] {
      if smallestIndex == -1 {
        smallestIndex = i
        continue
      }
    }
    if ele == sorted[1] {
      if middleIndex == -1 {
        middleIndex = i
        continue
      }
    }
    if ele == sorted[2] {
      if largestIndex == -1 {
        largestIndex = i
        continue
      }
    }
  }

	return smallestIndex, middleIndex, largestIndex
}

package main

// finds the smallest, middleIndex, and largestIndex answers
func IndexOfSmallMiddleLarge(arr []float64) (int, int, int){
  var smallestIndex, middleIndex, largestIndex int = 0,1,2
  var  middleValue, largestValue float64 = 0,0
  for index, ele := range arr {
    if ele > largestValue {
      middleValue = largestValue
      largestValue = ele
      largestIndex = index
    } else if ele > middleValue {
      middleValue = ele
      middleIndex = index
    } else {
      smallestIndex = index
    }
  }
  return smallestIndex, middleIndex, largestIndex
}


package methods

type MethodResults struct {
  Name string
  // each index corresponds to the index of the answer
  Results []float64
}

var Methods []*MethodResults

func init() {
  Methods = []*MethodResults{googleMethod, wikipediaMethod, bingMethod}
}

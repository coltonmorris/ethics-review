+ **main**
  - area that can listen to file changes fires off methods.
  - calculates methods weights for quorum while waiting for next file change
```
// initially calculate weights
let weights[] = calculateMethodWeights()

onFileChange(() => {
  results = runQuorum(weights)
  pretty-print results
  weights = calculateMethodWeights()
  saveQuorumResults()
})

```

+ **quorum**
  - area that listens to all methods and calculates the average answer from all the methods that have currently been returned.
  - pretty print results
```
```

+ **methods** each method records it's results to a csv
  - Saving CSV fields: (question, answer1,  answer2,  answer3, prediction1, prediction2, prediction3, correctAnswer, timer, more?)
```
```

+ **Redux**
  - wrap store to make middleware. separate store into separate file
    - Make a middleware that passes actions to a function (epic)
      + Epic for starting methods
      + Quorum Epic that runs whenever a methodFinished. 
  - combine reducers is just a reducer that splits the state and passes it to the reducers.
```
state = {
  history
  question
  answers

  method1
  method2
  method...
}
```

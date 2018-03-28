import _ from 'lodash'
import { Promise } from 'bluebird'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators';
import util from 'util'
import methods from './methods'
import colors from 'colors'


// formats the quorums result to whatever we want
let saveQuorumResults = (methods, question, answers, finalGuess) => ({
  question: question,
  answers: answers,
  methods: methods,
  finalGuess: finalGuess,
})


// TODO idea for giving weights for each method:
//    have each method record it's guesses into a csv, with the actual correct guess.
//    calculate how often the method is correct by parsing the csv file
//    methods that are more accurate are given more credit

// calculates the odds for each answer, ending with a final guess
let evaluateQuorum = (methods, answers) => {
  // Populate answerWeights object
  // answerWeights = { 'answer1': [0.3, ...] ]
  let answerWeights = {}
  _.map(answers, (answer) => {
    answerWeights[answer] = { weights: [] }
  })
  _.map(methods, (method) => {
    _.mapValues(_.omit(method, ['method']), (value) => {
      answerWeights[value.answer].weights.push(value.weight) 
    })
  })

  // this shit is ugly. was tired and didn't bother making it pretty...
  // calculate average and save it to finalGuess
  let finalGuess = {
    smallest: {},
    middle:   {},
    largest:  {},
  }
  let smallest = 0;
  let middle = 0;
  let largest = 0;
  _.mapKeys(answerWeights, (value, key) => {
    let count = 0 
    let sum = 0
    _.map(answerWeights[key].weights, (weight) => {
      sum += weight
      count++
    })

    let average = sum / count
    answerWeights[key].average = average
    if (average > largest) {
      largest = average
      finalGuess.smallest = finalGuess.middle
      finalGuess.middle = finalGuess.largest
      finalGuess.largest = { answer: key, averageWeight: average }
    } else if (average > middle) {
      middle = average
      finalGuess.smallest = finalGuess.middle
      finalGuess.middle = { answer: key, averageWeight: average }
    } else {
      smallest = average
      finalGuess.smallest = { answer: key, averageWeight: average }
    }
  })

  return finalGuess
}

export default (question, answers) => {
  // pass each method the question and answers. They will return an object. Save that object 

  // TODO make a method that just sleeps 10 seconds
  // TODO implement the rxjs method  takeUntilWithTime to finish the quorum before the 10 seconds is up
  return Observable.combineLatest(..._.map(methods, (method) => (
    Observable.from(method(question, answers))
  ))).map(methods => {
    // just received a new result from a method
    let finalGuess = evaluateQuorum(methods, answers)
    let quorumResult = saveQuorumResults(methods, question, answers, finalGuess)
    printQuorumResult(quorumResult)
    return quorumResult
  })
}

let printQuorumResult = (result) => {
  // console.log("Quorum Result:\n", util.inspect(quorumResult, false, null))
  console.log('\t\t\tQuorum Result:'.black)


  _.mapKeys(result, (value, key) => {
    if (key == 'finalGuess') {
      _.mapKeys(value, (innerValue, innerKey) => {
        console.log('yo: ', innerValue, innerKey)
        if (innerKey == 'smallest') console.log(innerKey.red, innerValue.red)
        if (innerKey == 'middle') console.log(innerKey.yellow, innerValue.yellow)
        if (innerKey == 'largest') console.log(innerKey.green, innerValue.green)
      })
    } else {
      console.log(key.black, util.inspect(result[key],false,null))
    }
  })
}

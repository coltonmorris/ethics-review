import _ from 'lodash'
import { Promise } from 'bluebird'
import { Observable } from 'rxjs/Rx'
import { mergeMap } from 'rxjs/operators';


import methods from './methods'

let quorumResults = {}
export { quorumResults as quorumResults }
// quorumResults is an object like this: {
//   exampleResultMethod: {
//     smallest: {answer: 'one', weight: .3},
//     middle: {answer: 'two', weight: .1},
//     largest: {answer: 'three', weight: .6},
//   }
// }

export default (question, answers) => {
  mainQuorum(question, answers)
  // Observable.from(mainQuorum(question, answers))
  //   .take(3)
  //   .subscribe((x) => {
  //     console.log('in start quorum: ', x)
  //     console.log('in start quorum: ', x)
  //     console.log('in start quorum: ', x)
  //   })
}

let saveResult = (result) => {
  quorumResults[result.method] = _.omit(result, ['method'])
}

let evaluateQuorum = () => {
  console.log('-- Quorum Evaluation --')

  // evaluate each result's largest weights
  _.mapKeys(quorumResults, (value, key) => {
    console.log('key: ', key)
    console.log('Smallest: ', value.smallest)
    console.log('Middle: ', value.middle)
    console.log('Largest: ', value.largest)
  })

  console.log('-----------------------')
}

function mainQuorum(question, answers) {
  // pass each method the question and answers. They will return an object. Save that object 

  // TODO make a method that just sleeps 10 seconds
  // TODO implement the rxjs method  takeUntilWithTime to finish the quorum before the 10 seconds is up
  // TODO look into using combineLatest
  console.log('methods: ', methods)
  Observable.from(_.map(methods, (method) => (
    // TODO from?
    // Observable.of(method(question, answers))
    method(question, answers)
    // let methodResult = Observable.from(method(question, answers))
    // methodResult.subscribe(result => {
    //   saveResult(result)
    //   evaluateQuorum()
    // })
    // return methodResult
  )))
  // TODO START RIGHT HERE. merge? the results of the lodash map
  //    and whenever one of those methods emits from it's observer,
  //    saveResult, and evaluate.
    .mergeMap(val => Observable.of('hi'))
    .subscribe((x) => {
      console.log('in main quorum: ', x)
  })
}

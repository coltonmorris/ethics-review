import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators'
import puppeteer from 'puppeteer'


export default (question, answers) => {
  return Observable.forkJoin(
    ..._.map(answers, (answer) => (
      Observable.from(getResults(question, answer))
    ))).map(results => {
      return transfromResults(results)
    })
}

let getResults =  async (question, answer) => {
  return { answer: answer, count: 0 }
};

let transfromResults = (results) => {
  results.sort((a,b) => {
    a = parseInt(a.count)
    b = parseInt(b.count)
    if (a < b)
      return -1;
    if (a > b)
      return 1;
    return 0;
  })

  // check if count is 0
  results = _.map(results, (result) => {
    if (result.count == 0) return { ...result, count: 1 }
    return result
  })

  let sum = _.reduce(results, (val, ele) => (val + parseInt(ele.count)), 0)

  let final = {
    smallest: { answer: results[0].answer, weight: results[0].count/sum },
    middle: { answer: results[1].answer, weight: results[1].count/sum },
    largest: { answer: results[results.length-1].answer, weight: results[results.length-1].count/sum },
    method: 'wikipedia',
  }
  return final
}

import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators'
import axios from 'axios'
let WordPOS = require('wordpos')

export default (question, answers) => {
  return Observable.from(getNouns(question))
    .map(nouns => {
      return Observable.forkJoin(
        ..._.map(answers, (answer) => {
          return Observable.from(getBody(nouns, answer))
        })
      ).map(results => {
          return transfromResults(results)
      })
    })
}

let getNouns = async(question) => {
  let wordpos = new WordPOS()
  return await wordpos.getNouns(question, nouns => nouns)
}

const tally = (str, ans) => {
  let re = new RegExp(ans, "gi")
  let result = str.match(re)
  if(!result){ return 0 }
  else { return result.length }
}

let getBody = async(questionNouns, answer) => {
  let transAnswer = _.replace(_.escape(answer), " ", "%20")
  let url = `https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=${transAnswer}`

  let result = await axios.get(url)
  const pageData = result.data.query.pages[Object.keys(result.data.query.pages)[0]].extract

  let total = 0
  questionNouns.forEach((noun) => {
    total += tally(pageData, noun)
  })

  return { answer: answer, count: total }
}

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

  let sum = _.reduce(results, (val, ele) => (val + parseInt(ele.count)), 0)

  let final = {
    smallest: { answer: results[0].answer, weight: 0.01 },
    middle: { answer: results[1].answer, weight: 0.09 },
    largest: { answer: results[results.length-1].answer, weight: 0.9 },
    method: 'wikipedia',
  }
  return final
}

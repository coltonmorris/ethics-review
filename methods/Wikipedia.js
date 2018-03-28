import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators'
import axios from 'axios'
let WordPOS = require('wordpos')

export default (question, answers) => {
  return Observable.from(getNouns(question))
    .map(nouns => {
      console.log('~~nouns', nouns) //ok
      return Observable.forkJoin(_.map(answers, (answer) => {
          console.log('iter', answer)
          return Observable.from(getBody(nouns, answer))
          .map((val) => {
            console.log('^^^^^^^', val)
            return val
          })
      })).map(results => {
          return transfromResults(results)
        })
    })
}

let getNouns = async(question) => {
  let wordpos = new WordPOS()
  return await wordpos.getNouns(question, nouns => nouns)
}

const count = (str, ans) => {
  let re = new RegExp(ans, "gi")
  return str.match(re).length
}

let getBody = async(questionNouns, answer) => {
  let transAnswer = _.replace(_.escape(answer), " ", "%20")
  let url = `https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=${transAnswer}`

  console.log('before awiat')
  return await axios.get(url)
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

  let sum = _.reduce(results, (val, ele) => (val + parseInt(ele.count)), 0)

  let final = {
    smallest: { answer: results[0].answer, weight: 0.01 },
    middle: { answer: results[1].answer, weight: 0.09 },
    largest: { answer: results[results.length-1].answer, weight: 0.9 },
    method: 'wikipedia',
  }
  return final
}

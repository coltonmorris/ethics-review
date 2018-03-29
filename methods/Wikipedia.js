import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators'
import axios from 'axios'
import fs from 'fs'
let WordPOS = require('wordpos')

export default (question, answers) => {
  return Observable.from(getNouns(question).then((nouns) => {
      return Promise.map(answers, (answer) => {
        return getBody(nouns, answer)
      })
      .then((res) => transfromResults(res))
    }))
}
  // return getNouns(question).then((nouns) => {
  //   return Observable.forkJoin(
  //       ..._.map(answers, (answer) => {
  //         return Observable.from(getBody(nouns, answer))
  //       })
  //     ).map(results => {
  //         return transfromResults(results)
  //     })
  //   })
// }

let getNouns = async(question) => {
  let wordpos = new WordPOS()
  return await wordpos.getNouns(question, nouns => nouns)
}

function occurrences(string, subString, allowOverlapping, filename) {
  console.log('searching substring', subString)
  fs.writeFile(filename+'.txt', string, function (err) {
    if (err) throw err;
  })
  string += "";
  subString += "";
  if (subString.length <= 0) return (string.length + 1);

  var n = 0,
    pos = 0,
    step = allowOverlapping ? 1 : subString.length;

  while (true) {
    pos = string.indexOf(subString, pos);
    if (pos >= 0) {
      ++n;
      pos += step;
    } else break;
  }
  return n;
}

let getBody = async(questionNouns, answer) => {
  let transAnswer = _.replace(_.escape(answer), " ", "%20")
  let url = `https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=${transAnswer}`

  let result = await axios.get(url)
  const pageData = result.data.query.pages[Object.keys(result.data.query.pages)[0]]
  let body
  if(pageData.hasOwnProperty('missing')){
    body = ""
  } else {
    body = pageData.extract
  }

  let total = 0
  questionNouns.forEach((noun) => {
    console.log('noun', noun, occurrences(body, noun, false))
    total += occurrences(body, noun, false, answer)
  })

  return { answer: answer, count: total }
}

let transfromResults = (results) => {
  console.log('hckjlekrj', results)
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

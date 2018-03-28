import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators'
import puppeteer from 'puppeteer'
import axios from 'axios'

let answers = ['yorkshire pudding', 'cabbage patch', 'south america']

let getBody = async(answer) => {
  let questionNouns = ['breakfast', 'england']
  let transAnswer = _.replace(_.escape(answer), " ", "%20")
  let url = `https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=${transAnswer}`

  console.log('before awiat')
  return await axios.get(url)
};

let getNouns = async(question) => {
  let wordpos = new WordPOS()
  return await wordpos.getNouns(question, nouns => nouns)
}

Observable.from(getNouns(question))
  .map(nouns => {
    // console.log('~~nouns', nouns) //ok
    return Observable.forkJoin(_.map(answers, (answer) => {
        // console.log('iter', answer)
        return Observable.from(getBody(nouns, answer))
    })).map(results => {
        return transfromResults(results)
      })
  })

Observable.from(getBody()).map((res) => {
  console.log('hi', Object.keys(res))
  return res
})

import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap } from 'rxjs/operators';
const puppeteer = require('puppeteer')


export default async (question, answers) => {
  // TODO make it an observable that takes 3. use yield instead of return
  // var result = Rx.Observable.from(iterator).take(3);
  // result.subscribe(x => console.log(x));

  // TODO start here
  return Observable.of(_.map(answers, async (answer) => {
    return Observable.of(getResults(question, answer)
  })).subscribe((res) => {
    console.log('inside number of google results')
    res.sort((a,b) => {
      a = parseInt(a.count)
      b = parseInt(b.count)
      if (a < b)
        return -1;
      if (a > b)
        return 1;
      return 0;
    })

    let sum = _.reduce(res, (val, ele) => (val + parseInt(ele.count)), 0)

    let final = {
      smallest: { answer: res[0].answer, weight: res[0].count/sum },
      middle: { answer: res[1].answer, weight: res[1].count/sum },
      largest: { answer: res[res.length-1].answer, weight: res[res.length-1].count/sum },
      method: 'numberOfGoogleResults',
    }
    console.log('final google results: ', final)
    return final
  })
}

let getResults =  async (question, answer) => {
  let queryUrl = `https://www.google.com/search?q=${question} ${answer}`
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto(queryUrl);

  const resultHandle = await page.$('#resultStats')
  let count = await page.evaluate(result => result.innerHTML, resultHandle)

  count = count
    .split(' ')[1]
    .replace(/,/g,'')


  browser.close();
  return { answer: answer, count: count}
};


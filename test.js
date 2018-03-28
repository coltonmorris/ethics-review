import { Promise } from 'bluebird'
import _ from 'lodash'
import { Observable } from 'rxjs/Rx'
import { mergeMap, map } from 'rxjs/operators'
import puppeteer from 'puppeteer'

let main = async() => {
  let question = 'why though?'
  let answer = 'cuz'
  let queryUrl = `https://www.bing.com/search?q=${question} ${answer}`
  const browser = await puppeteer.launch()
  const page = await browser.newPage()
  await page.goto(queryUrl)

  const resultsSelector = '.sb_count'
  await page.waitForSelector(resultsSelector)

 // Extract the results from the page.
  let count = await page.evaluate(resultsSelector => {
    return document.querySelector(resultsSelector).innerHTML
  }, resultsSelector)

  return count.split(' ')[0].replace(/,/g,'')
}

// main().then((res) => {
Observable.from(main()).map((res) => {
  console.log('hi')
  return res
}).subscribe((res) => {
  console.log('subscribe')
  console.log(res)
})

const puppeteer = require('puppeteer')
const fs = require('fs') 
const path = require('path')    
const Promise = require('bluebird')
const _ = require('lodash')


const filePath = path.join(__dirname, 'ocr_output.txt')
let question = ""
let answers = []


let data = fs.readFileSync(filePath, 'utf-8')

data = data.trim().split('\n')
console.log('Read from file: \n', data)
// remove first newline
let endOfQuestionIndex = data.findIndex((line) => line === '' )
question = data.slice(0, endOfQuestionIndex).join(' ')
answers = data.slice(endOfQuestionIndex+1, data.length)
answers = _.compact(answers)

// tODO strip white space from answers array

console.log('************ PARSED **************')
console.log('answers: ', answers)
console.log('question: ', question)



let getResults = async (question, answer) => {
  let queryUrl = `https://www.google.com/search?q=${question} ${answer}`
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto(queryUrl);

  const resultHandle = await page.$('#resultStats')
  let html = await page.evaluate(result => result.innerHTML, resultHandle)

  html = html
    .split(' ')[1]
    .replace(/,/g,'')


  browser.close();
  return { answer: answer, html: html}
};

function compareResult(a,b) {
  a = parseInt(a.html)
  b = parseInt(b.html)
  if (a < b)
    return -1;
  if (a > b)
    return 1;
  return 0;
}

Promise.map(answers, async (answer) => {
  return getResults(question, answer)
}).then((results) => {
  /* do something here with the finished results */
  results.sort(compareResult)
  console.log('***********************************')
  console.log('************ Results **************')
  console.log('***********************************')
  console.log(results)
  console.log('Smallest: ', results[0])
  console.log('Largest: ', results[results.length-1])
})


import startQuorum from './quorum'
import { Observable } from 'rxjs/Rx'
import _ from 'lodash'
import path from 'path'
import fs from 'fs'
import util from 'util'
import Papa from 'papaparse'

import methods from './methods'


const ocrOutputFile = 'ocr_output.txt'
console.log(`Listening to changes made by tessaract to: ${ocrOutputFile}\n`)

let parseOcr = () => {
  const filePath = path.join(__dirname, ocrOutputFile)
  let question = ""
  let answers = []
  let data = fs.readFileSync(filePath, 'utf-8')
  console.log('Read from file: \n', data)

  data = data.trim()
  let i = data.indexOf('?')
  question = data.substring(0,i)
  question = question.replace(/\n/g,' ').replace(/'/g,'"')

  answers = data.substring(i+1,data.length).trim().split('\n')
  answers = _.compact(answers)

  return [question, answers]
}

// Watch the file output by tessaract
fs.watchFile(ocrOutputFile, (curr, prev) => {
  console.log('**********************************')
  console.log('**********************************')
  console.log('**********************************')

  let ocr = parseOcr()
  let question = ocr[0]
  let answers = ocr[1]


  console.log('**********************************')
  console.log('**********************************')
  console.log('**********************************')

  startQuorum(question, answers).subscribe((quorumResult) => {
    // TODO have each method record it's own guesses
    // TODO have the quorum record it's own guesses
    // TODO have a user input to select the correct answer for csv's
  
    // TODO always save the same method to the corresponding field 
    // fields: [question, answer1, answer2, answer3, correctAnswer, guessedAnswer, methodname_0, methodweight_0, methodname_1, methodweight_1, methodname_2, methodweight_2, methodname_3, methodweight_3, methodname_4, methodweight_4, methodname_5, methodweight_5, methodname_6, methodweight_6, methodname_7, methodweight_7, methodname_8, methodweight_8, methodname_9, methodweight_9, methodname_9, methodweight_9]

    // TODO this needs to only append a line rather than add a header each time
    // save to csv
    let csv = Papa.unparse({
      fields: ['question', 'answers', 'correctAnswer'],
      data: [question, answers, 'fill me in']
    })
    fs.appendFile('quorumHistory.csv', csv, function (err) {
      if (err) throw err;
    })
  })
})

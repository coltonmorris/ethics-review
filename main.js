import startQuorum, {quorumResults} from './quorum'
import { Observable } from 'rxjs/Rx'
import _ from 'lodash'
import path from 'path'
import fs from 'fs'

import methods from './methods'

const ocrOutputFile = 'ocr_output.txt'

// Watch the file output by tessaract
fs.watchFile(ocrOutputFile, (curr, prev) => {
  const filePath = path.join(__dirname, ocrOutputFile)
  let question = ""
  let answers = []

  let data = fs.readFileSync(filePath, 'utf-8')

  data = data.trim().split('\n')
  console.log('Read from file: \n', data)
  // remove first newline, which delimits the question
  let endOfQuestionIndex = data.findIndex((line) => line === '' )
  question = data.slice(0, endOfQuestionIndex).join(' ')

  answers = data.slice(endOfQuestionIndex+1, data.length)
  answers = _.compact(answers)

  console.log('**********************************')
  console.log('************ PARSED **************')
  console.log('**********************************')
  console.log('answers: ', answers)
  console.log('question: ', question)

  // get the result of this quorum, and save it to a variable. Have an observer watching that variable that uploads the csv
  startQuorum(question, answers)
})


// Observable.of(methodsObserver)
//   .subscribe(x => {
//     console.log(x)
//    // record csv
//    //    use google cloud storage?
//   })




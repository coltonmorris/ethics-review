# Contributing

Info on code base and contributing to it.

## Adding new Methods
To add a new method, you must fulfill 3 requirements:
1. Make a function that fulfills finding the appropriate weights for the 3 questions, and returns on object of type `*MethodResults`.
2. Make a function that fulfills the StartFunc interface. The function declaration should look something like this: `var startNewMethod StartFunc = func(qna *QandA, doneChannel chan *MethodResults) {`
3. Append this start function to the `StartMethods` array in `methods/main.go`.


## guess-it-1

This is an excercise in the 01-edu curriculum about guessing in what range a next number might fit based on previous numbers.

It uses statistics related functions from an earlier exercise.

\
run with:
``` bash
docker compose up
```
\
Guess against at least these three ai models: "average", "big-range" and "median".\
Also vailable: "correlation-coef", "huge-range", "linear-regr", "mse" and "nic" 
```
http://localhost:3000/?guesser=average
```

\
To make testing work:
- Abide by the instructed folder structure
- script.sh should include these lines:\
cd student\
go run solution.go  
- Make sure your project go version is 1.23
- After a failed attempt, delete container and image from docker desktop

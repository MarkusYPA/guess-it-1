## guess-it-1

Guess-it-1 is an excersice in the 01-edu curriculum and a program that guesses in what range a next number might fit based on previous numbers and statistical calculations about the median and standard deviation.

To run the auditing program, extract the provided zip file and copy the student folder to what is by default called guess-it-dockerized. 

Now run the testing program with:
``` bash
docker compose up
```

Guess against at least these three ai models: "average", "big-range" and "median".\
Also available: "correlation-coef", "huge-range", "linear-regr", "mse" and "nic" 

To test "average", navigate to this URL:
```
http://localhost:3000/?guesser=average
```

### Folder structure

Guess-it-1 contains five programs. The main program is in the "student" folder, the others are there to help analyze the given data sets and to test how the main program fares against ai models.

All helper programs are meant to be compiled into executables and run in the student folder. They are:

#### 1. Freqs

Freqs prints out the frequencies of all non-outlier values. This information can be used to create a static range guess rather successfully if not complitely observing the spirit of the excercise.

#### 2. Numwriter

Numwriter writes values to standard output from a data set given as arguments (e.g. ./numwriter 2 3) like the testing program does. This output can be piped as input to the main program.

#### 3. SD size tester

SD size tester tests which values for the standard deviation multiplier produce the most points by running all the 15 data sets with Numwriter to the main program multiple times, giving different values as arguments to the main program from a user edited slice. It prints out the average points yield for each value.

#### 4. Run the Numbers

Run the Numbers tests the main program against known results from ai models, printing the points yield from the program for each data set and a W or an L to signify a win or a loss against the "average" and the "median" models. This is much faster than running the provided tester program over and over.


  
### Notes about getting testing to work
To make testing work:
- Abide by the instructed folder structure
- script.sh should include these lines:\
cd student\
go run solution.go  
- Make sure your project go version is 1.23 (or the same the dockerized program uses)
- To run the tester with a newer version of the main program: close Docker, delete the container and the image from Docker desktop, copy the updated solution.go over the old one and run "docker compose up" again.


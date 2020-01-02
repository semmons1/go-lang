 # Command Line Quiz Game 
 This simple program can be run by changing to the directory in which these files have been cloned. There are two flags which should be considered when running `go run main.go`.<br/> Quiz questions and answers are stored in a separate CSV file, and are read in one at a time to the user.<br/> The user should answer each question with an integer input, such as "`10`,`20`, etc. 
 `-time` : This flag sets an argument which translates into a timer for program expiration that defaults to thirty seconds. Once this timer expires, the entire program will exit, regardless of where the user is. An example on how to change this flag is `-time = 40`. This will set the timer to forty seconds.<br/>
 `-file` : This flag sets an argument that specifies the name of your CSV file which contains the questions and answers for this quiz. This flag defaults to "quizfile.csv". This file is encouraged to be modified, and if renamed, be sure to include `-file = "yourfilename.csv"` when executing `go run main.go`. 
 
 

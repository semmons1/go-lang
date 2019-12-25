package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var begin string

	/* To ensure that user is ready to go, parse input in main, and then call quiz function. That function
	contains the CSV reader utility, along with an encompassing timer channel that will dictate when
	the program exits if all questions are not answered.
	*/

	fmt.Printf("Are you ready to start? (y/n)\n")
	fmt.Scanln(&begin)
	length := flag.Int("time", 30, "Flag for specifying timer for the quiz, in seconds.")
	fName := flag.String("file", "quizFile.csv", "Flag for specifying the CSV containing quiz questions.")
	flag.Parse()

	if begin == "y" {
		takeQuiz(*length, *fName)
	} else {
		os.Exit(0)
	}
}

func takeQuiz(length int, fName string) {
	var correctCount int32
	var counter int32
	var ans string

	/* As long as the CSV file is in the same directory as "quiz.go", the os and CSV packages will
	track it down, and open it for reading. If you want to change the filename, be sure to change
	the name here as well.
	*/

	quizFile, _ := os.Open(fName)
	reader := csv.NewReader(bufio.NewReader(quizFile))
	timer := time.NewTimer(time.Second * time.Duration(length))

	go func() {
		for {
			line, error := reader.Read()

			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}

			fmt.Printf("What does %s equal?\n", line[0])
			fmt.Scanln(&ans)

			if ans == line[1] {
				correctCount++
			}

			counter++
		}
		fmt.Printf("Nice job, you answered %d out of %d questions correctly.\n", correctCount, counter)
		os.Exit(1)
	}()

	select {
	case <-timer.C:
		fmt.Printf("\nThe timer has expired.\n")
		fmt.Printf("You were able to answer %d out of %d questions correctly.\n", correctCount, counter)
	}
}

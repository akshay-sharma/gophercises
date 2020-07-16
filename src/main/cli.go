package main

import (
	"flag"
	"time"
	"quiz"
)

func main() {
	filePath := flag.String("file", "/home/akshay/repos/gophercises/src/quiz/problems.csv", "Path of file that has question set")
	timeDuration := flag.Int("time", 10, "Duration of the quiz")
	flag.Parse()
	quiz.Handler(*filePath, time.Duration(*timeDuration) * time.Second)
}

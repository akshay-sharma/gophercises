package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func Handler(filePath string, duration time.Duration) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		exitMsg(fmt.Sprintln("Could not open the csv file ", filePath, err))
	}
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		exitMsg(fmt.Sprintln("Could not parse csv file ", err))
	}
	problems := parseQuestions(records)
	score := 0
	timerChannel := time.After(duration)
	userInputChannel := make(chan string)

	for _, problem := range problems {
		fmt.Println(problem.question)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			userInputChannel <- ans
		}()
		select {
		case <-timerChannel:
			fmt.Println("Time elapsed")
			fmt.Println(score, "/", len(records))
			return
		case input := <-userInputChannel:
			if strings.TrimSpace(input) == problem.answer {
				score++
			}
			break
		}
	}
	fmt.Println(score, "/", len(records))
}

func parseQuestions( records [][] string) []problem {
	parsedProblems := []problem{}
	for _, record := range records {
		parsedProblems = append(parsedProblems, problem{
			question: strings.TrimSpace(record[0]),
			answer:   strings.TrimSpace(record[1]),
		})
	}
	return parsedProblems
}

func exitMsg(message string) {
	fmt.Println(message)
	os.Exit(1 )
}

type problem struct {
	question string
	answer string
}

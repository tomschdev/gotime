package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type T string
type QA[R, S any] struct {
	question R
	answer   S
}
type Quiz struct {
	sequence []QA[T, T]
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv formatted questions with each line containing: question, answer")
	timeLimit := flag.Int("time", 5, "time limit for quiz")
	flag.Parse()

	quiz := readQuiz(*csvFilename)
	correct := 0
	total := len(quiz.sequence)
	quiz_ch, res_ch := make(chan QA[T, T], total), make(chan int)
	var res int

	//load all quiz problems into buffered channel
	for _, v := range quiz.sequence {
		quiz_ch <- v
	}

	//receive and process all quiz problems on new goroutine
	go runQuiz(quiz_ch, res_ch, time.Duration(*timeLimit)*time.Second)

	//process answers and results of quiz
	for {
		res = <-res_ch
		if res == 1 {
			correct += res
		}
		if res == 2 {
			fmt.Println("time's up!")
			break
		}
		if res == 3 {
			fmt.Println("quiz complete!")
			break
		}
	}
	fmt.Printf("you scored: %v/%v\n", correct, total)
}
func runQuiz(quiz_ch chan QA[T, T], res_ch chan int, dur time.Duration) {
	input := make(chan string)
	go func(in chan<- string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			handleErr(err)
			in <- text
		}
	}(input)

	timeup := time.After(dur)
	for q := range quiz_ch {
		fmt.Println("Question: ", q.question)
		select {
		case text := <-input:
			ans := strings.Trim(strings.Trim(text, " "), "\n")
			// fmt.Printf("%s %s %v\n", ans, q.answer, (T(ans) == q.answer))
			if T(ans) == q.answer {
				res_ch <- 1
			}
		case <-timeup:
			res_ch <- 2
			return
		}
	}
	res_ch <- 3
}
func readQuiz(fileName string) Quiz {
	var f *os.File
	var err error
	f, err = os.Open(fileName)
	handleErr(err)
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	handleErr(err)
	return createQuiz(data)
}
func createQuiz(data [][]string) Quiz {
	var quiz Quiz
	var qa QA[T, T]
	quiz.sequence = make([]QA[T, T], 0)
	for _, v := range data {
		qa = QA[T, T]{T(strings.TrimSpace(v[0])), T(strings.TrimSpace(v[1]))}
		quiz.sequence = append(quiz.sequence, qa)
	}
	return quiz
}
func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

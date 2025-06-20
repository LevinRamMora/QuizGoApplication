package Service

import (
	"QuizGoApplication/Model"
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var result int
var wrongCount int
var wrongAnswers = map[string]string{}
var ofTAnswers = map[string]string{}
var invalidInputs = map[string]string{}

func StartQuiz(questions *[]Model.Questions, name *string) {
	var duplicates = []int{}

	i := 1
	for i <= 10 {
		if wrongCount < 3 {
			qNumber := rand.Intn(len(*questions))
			if slices.Contains(duplicates, qNumber) {
				continue
			}
			start := time.Now()
			fmt.Println("\n" + strconv.Itoa(i) + ". " + (*questions)[qNumber].Question + "    Time: " + start.Format("02-01-2006 15:04:05"))
			for _, option := range (*questions)[qNumber].Options {
				fmt.Println(option)
			}
			getInputAndCalculate(&(*questions)[qNumber].Answer, &result, &wrongCount, &start, &(*questions)[qNumber])

			i++
			duplicates = append(duplicates, qNumber)
		} else {
			break
		}
	}
	evaluateAndDisplay(&result, name)

}

func getInputAndCalculate(ans *string, result *int, wrongCount *int, start *time.Time, question *Model.Questions) {

	for i := 1; i <= 2; i++ {

		fmt.Print("Answer: ")
		rawData := bufio.NewReader(os.Stdin)
		str, err := rawData.ReadString('\n')
		if err != nil {
			fmt.Println(errors.New("something went wrong"))
		}
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.TrimSpace(str)
		end := time.Now()
		ch := str[0]
		fmt.Printf("Submitted with %v, Time: %v \n", str, end.Format("02-01-2006 15:04:05"))
		difference := end.Sub(*start)
		if difference <= 60*time.Second {
			if (ch >= 65 && ch <= 68) || (ch >= 97 && ch <= 100) {
				if strings.EqualFold(string(ch), *ans) {
					*result = *result + 10
					break
				} else {
					*wrongCount++
					if question.Answer == "A" {
						wrongAnswers[question.Question] = question.Options[0]
					} else if question.Answer == "B" {
						wrongAnswers[question.Question] = question.Options[1]
					} else if question.Answer == "C" {
						wrongAnswers[question.Question] = question.Options[2]
					} else if question.Answer == "D" {
						wrongAnswers[question.Question] = question.Options[3]
					}
					break
				}
			} else {
				if i == 2 {
					fmt.Print("Invalid input \n")
					if question.Answer == "A" {
						invalidInputs[question.Question] = question.Options[0]
					} else if question.Answer == "B" {
						invalidInputs[question.Question] = question.Options[1]
					} else if question.Answer == "C" {
						invalidInputs[question.Question] = question.Options[2]
					} else if question.Answer == "D" {
						invalidInputs[question.Question] = question.Options[3]
					}
				} else {
					fmt.Print("Invalid input, Enter Correct ")
				}

				continue
			}
		} else {
			fmt.Println("Time's Up, Answer is not considered.")
			if question.Answer == "A" {
				ofTAnswers[question.Question] = question.Options[0]
			} else if question.Answer == "B" {
				ofTAnswers[question.Question] = question.Options[1]
			} else if question.Answer == "C" {
				ofTAnswers[question.Question] = question.Options[2]
			} else if question.Answer == "D" {
				ofTAnswers[question.Question] = question.Options[3]
			}
			break
		}
	}

}

func evaluateAndDisplay(result *int, name *string) {

	fmt.Println("Submitted Correct Answers: ", *result/10)
	fmt.Printf("*******************************************************\nResult: %v Marks out of 100\n", *result)
	if *result < 70 {
		fmt.Printf("Final Verdict: Fail \nBetter luck next time %v.\n*******************************************************", *name)
	} else {
		fmt.Printf("Final Verdict: Pass \nCongratulations %v!!\n*******************************************************", *name)
	}
	if len(wrongAnswers) > 0 {
		index := 1
		fmt.Println("\n*******************************************************\nWrong Answers: ")
		for question, answer := range wrongAnswers {
			fmt.Printf("%v. %v\nAnswer: %s\n", index, question, answer)
			index++
		}
	}
	if len(ofTAnswers) > 0 {
		index := 1
		fmt.Println("\n*******************************************************\nOut of Time Answers:")
		for question, answer := range ofTAnswers {
			fmt.Printf("%v. %v\nAnswer: %s\n", index, question, answer)
			index++
		}
	}
	if len(invalidInputs) > 0 {
		index := 1
		fmt.Println("\n*******************************************************\nInvalid Input Answers:")
		for question, answer := range invalidInputs {
			fmt.Printf("%v. %v\nAnswer: %s\n", index, question, answer)
			index++
		}
	}
}

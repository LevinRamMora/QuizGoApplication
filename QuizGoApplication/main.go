package main

import (
	"QuizGoApplication/Model"
	"QuizGoApplication/Service"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var input string
	var name string
	QuestionsData, err := load()
	if err != nil {

		fmt.Println("Error Occurred", err)
	}
	fmt.Println("Welcome to the Quiz \nPlease enter your name: ")
	fmt.Scanln(&name)

	fmt.Printf("\nRead Instructions carefully that are mentioned below: \n" +
		"1. Quiz Contains 10 Questions. \n" +
		"2. Each Answer should be submit within 60 Seconds, if not Answer is not considered. \n" +
		"3. User got Terminated if user answered 3 wrong answers. \n" +
		"4. User Must provide single alphabet only (A/B/C/D) or (a/b/c/d).\n" +
		"5. User should score 70 Marks to pass.ð \n" +
		"All The Best. \n" +
		"Shall We Start? (Yes/No):  ")
	fmt.Scanln(&input)

	if strings.EqualFold("Yes", input) {
		Service.StartQuiz(&QuestionsData, &name)
	} else {
		fmt.Println("Thank You.")
	}

}

func load() ([]Model.Questions, error) {

	file, err := os.Open("data.json")
	if err != nil {
		return nil, errors.New("error when importing the Data")
	}
	defer file.Close()
	var data []Model.Questions
	rawData := json.NewDecoder(file)
	err2 := rawData.Decode(&data)
	if err2 != nil {
		return nil, errors.New("error when Decoding the Data")
	}
	return data, nil
}

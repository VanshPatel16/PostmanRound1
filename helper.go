package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type meal struct {
	Day      string
	Date     string
	MealName string
	Items    []string
}

// this function returns the index of the column of the day entered by the user
func getDayIndex(allCols [][]string, day string) int {
	ret := -1
	for i := 0; i < len(allCols); i++ {

		// fmt.Println(allCols[0][i])
		if day == allCols[i][0] {
			ret = i
		}
	}

	return ret

}

//this function returns the index of the row for the meal entered by the user

func getMealIndex(allCols [][]string, dayIndex int, meal string) int {

	for i := 0; i < len(allCols[dayIndex]); i++ {

		if allCols[dayIndex][i] == meal {
			return i
		}
	}

	return -1
}

// function number 1:
func mealNow(allCols [][]string, dayOfWeek string, mealOfTheDay string) {

	dayIndex := getDayIndex(allCols, dayOfWeek)                //this gave us the index of our day
	mealIndex := getMealIndex(allCols, dayIndex, mealOfTheDay) //this gave us the index of our meal

	fmt.Printf("\n\nThe following is the menu for %v's %v\n\n", dayOfWeek, allCols[dayIndex][mealIndex])
	// count := 0
	for i := mealIndex + 1; i < len(allCols[dayIndex]) && allCols[dayIndex][i] != dayOfWeek; i++ {
		// count++
		fmt.Printf("\t%v\n", allCols[dayIndex][i])

	}

	fmt.Println()

	return
}

// function number 2:
func itemCount(allCols [][]string, dayOfWeek string, mealOfTheDay string) int {

	dayIndex := getDayIndex(allCols, dayOfWeek)                //this gave us the index of our day
	mealIndex := getMealIndex(allCols, dayIndex, mealOfTheDay) //this gave us the index of our meal

	count := 0
	for i := mealIndex + 1; i < len(allCols[dayIndex]) && allCols[dayIndex][i] != dayOfWeek && allCols[dayIndex][i] != ""; i++ {
		count++

	} //this will loop through the Items in that meal on that day and the we return the item count

	return count

}

// function 3:
func findItemInMeal(allCols [][]string, dayOfWeek string, mealOfTheDay string, item string) int {

	dayIndex := getDayIndex(allCols, dayOfWeek)                //this gave us the index of our day
	mealIndex := getMealIndex(allCols, dayIndex, mealOfTheDay) //this gave us the index of our meal

	for i := mealIndex + 1; i < len(allCols[dayIndex]) && allCols[dayIndex][i] != dayOfWeek && allCols[dayIndex][i] != ""; i++ {

		if item == strings.ToUpper(allCols[dayIndex][i]) {

			return i

		}

	}

	return -1
}

//function 4 :

func menuToJson(allCols [][]string) {

	//we will store our menu into an array of our struct meal and then store that into json
	var fullMenu []meal
	for i := 0; i < len(allCols); i++ {

		for j := 2; j < len(allCols[i]) && allCols[i][j] != ""; {
			var currMeal meal
			var k int
			currMeal.Day = allCols[i][0]
			currMeal.Date = allCols[i][1]
			currMeal.MealName = allCols[i][j]

			for k = j + 1; k < len(allCols[i]) && allCols[i][k] != allCols[i][0] && allCols[i][k] != ""; k++ {
				currMeal.Items = append(currMeal.Items, allCols[i][k])
			}
			j = k + 1
			fullMenu = append(fullMenu, currMeal)
		}
	}

	dataForMyJson, err := json.MarshalIndent(fullMenu, "\t", "")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("mess.json", dataForMyJson, 0777)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("\njson file (mess.json) created succesfully in the same directory\n")
}

func jsonToStruct(fileName string) []meal {
	myFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("\nEnter proper name or make sure file is present\n")
	}

	//we will recieve the data in an array of structs
	var allMeals []meal
	err = json.Unmarshal(myFile, &allMeals)

	return allMeals

}

// func 5 : it is a reciever function that prints the data of the meal structure
func (m meal) printMeal() {
	fmt.Printf("Day : %v\n", m.Day)
	fmt.Printf("Date : %v\n", m.Date)
	fmt.Printf("MealName : %v\n", m.MealName)
	fmt.Printf("Items : \n%v\n", m.Items)
	fmt.Println()

}

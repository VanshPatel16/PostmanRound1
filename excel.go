package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	messMenuFile, err := excelize.OpenFile("Sample-Menu.xlsx") //opened the excel sheet
	if err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a day of the week :   ")
	inputScanner.Scan() //took input for the day of the week and made it capitalized as our menu is in capital
	dayOfWeek := strings.ToUpper(inputScanner.Text())

	fmt.Printf("Enter the meal : Breakfast,Lunch or Dinner :  ")
	inputScanner.Scan()
	mealOfTheDay := strings.ToUpper(inputScanner.Text()) //took meal of that day as input and made it capitalized as out menu is in capital

	// obtained the rows from the excel sheet
	allCols, err := messMenuFile.GetCols("Sheet1")

	if err != nil {
		fmt.Println("Error in fetching rows!")
		log.Fatal(err)
	} //checked for any errors

	mealNow(allCols, dayOfWeek, mealOfTheDay) //this function call prints the list of items for the mealOfTheDay ... FUNCTION 1

	numOfItems := itemCount(allCols, dayOfWeek, mealOfTheDay) //this function call assigns the number of items in that mealOfTheDay to the variable...FUNCTION @

	fmt.Printf("\nThere are %v items in %v's %v\n", numOfItems, dayOfWeek, mealOfTheDay) //we print the number of items and the mealOfTheDay

	fmt.Println("Enter the item you want to check in the menu :  ")
	inputScanner.Scan()
	item := strings.ToUpper(inputScanner.Text()) //took item as input that is to be searched and made it capitalized

	itemIndex := findItemInMeal(allCols, dayOfWeek, mealOfTheDay, item) //stored the index from our function call and displayed the result...FUNCTION 3

	if itemIndex >= 0 {
		fmt.Printf("\nThe item is Present\n\n")
	} else {
		fmt.Printf("Item not found!!")
	}

	fmt.Printf("\nDo you want to convert mess menu to json file ??Enter Y/N\n")
	inputScanner.Scan()
	res := inputScanner.Text()

	if res == "Y" {
		menuToJson(allCols) //calls the function to convert data to a json file

	} else {
		fmt.Printf("Okay json file will not be created\n")
	}

	fmt.Printf("\nDo you want to print each instance of the meals??Enter Y/N\n")
	inputScanner.Scan()
	res = inputScanner.Text()

	if res == "Y" {
		fmt.Printf("\nEnter the file name (eg. 'mess.json' in this case)\n")
		allMeals := jsonToStruct("mess.json")
		for i := 0; i < len(allMeals); i++ {
			allMeals[i].printMeal()
		}

	} else {
		fmt.Printf("Okay meals will not be printed\n")
	}

}

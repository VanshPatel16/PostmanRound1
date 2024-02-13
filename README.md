# PostmanRound1

Library used to manipulate excel sheet : Excelize

NOTE : The driver code could still be modified to have different inputs for each functions.But for now I have kept it so that the function 2 and 3 work on the same input as that of function 1.


Function 1 :

This function first uses 2 other functions to search the indices of the day and meal from the excel file which was taken as the input from the user.
Then it loops through the menu till the string same as day is encountered or if the column ends.

Function 2 :

Loops through the data after getting indices of day and meal while increasing a count.


Function 3 :
First gets the indices like the above 2 functions and then it searches linearly.

Function 4 :

First it loops through the excel and stores the data in an 'array of struct meal' (struct meal defined at the top) then that array is marshalled and saved as a json file named "mess.json" in the same directory

Function 5:

It unmarshalls the data from a json (the same json named mess.json created above in this case) and stores it in a array of structs.
then it prints all the instaces by looping through the slice/array with the use of a helper function "printMeal()"


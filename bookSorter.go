/*
Sometimes the output always fails/miss (maybe a bug), try to repeat a few times until you get the right output
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"sort"
)

func main() {

	// Declare our variables
	var bookListReader, returnedData string
	// End of declaration

	fmt.Println("----------------------------------------------")

	// Read and do validation for inputted number from user
	fmt.Println("Enter list of books code : ")

	for {

		scanner := bufio.NewScanner(os.Stdin) // bufio can be use for long string and string with spaces in it
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		bookListReader = scanner.Text() // Pass it to variable bookListReader

		process := processBook(bookListReader)

		if process == "err" {
			fmt.Println("Invalid inputted list")
		}else{

			returnedData = process
			break

		}

	}
	// End of input and validation


	fmt.Println("----------------------------------------------")
	fmt.Println("(INPUT) Book list :", bookListReader)
	fmt.Println("(OUTPUT) Sorted book list :", returnedData)
	fmt.Println("----------------------------------------------")

}

// Function for testing only
func testingBookSorter(s string) string {

	process := processBook(s)
	return process

}
// End of function

type ListBook struct {

	Category int
	Title string
	Size int

}

// Sorting function

type doMultiSort []ListBook

func (tempArray doMultiSort) Len() int{ 
	return len(tempArray) 
}

func (tempArray doMultiSort) Swap(i, j int){ 
	tempArray[i], tempArray[j] = tempArray[j], tempArray[i] 
}

func (tempArray doMultiSort) Less(i, j int) bool { 

	if tempArray[i].Category < tempArray[j].Category {
		return true
	}
	if tempArray[i].Category > tempArray[j].Category {
		return false
	}


	return tempArray[i].Size > tempArray[j].Size
}

// End of sorting function

// Phase 1 function, for arranging each data and then pass it to second phase in order to convert it to arranged string
func processBook(s string) string {

	var tempString string
	var tempArray []ListBook
	var arrC6, arrC7, arrC0, arrC9, arrC4, arrC8, arrC1, arrC2, arrC5, arrC3 []ListBook
	// var c6, c7, c0, c9, c4, c8, c1, c2, c5, c3 string

	words := strings.Split(s, " ") // Split the book list string into array

	// fmt.Println(len(words))
	// fmt.Println(words[1])

	for i := 0; i < len(words); i++ {

		rnBook := words[i]
		sLength := len(rnBook)

		if sLength > 3 && sLength < 5 {

			sCategory, _ := strconv.Atoi(bytesToString(rnBook[0]))
			sTitle := bytesToString(rnBook[1])

			combine := bytesToString(rnBook[2]) + bytesToString(rnBook[3])
			sSize, _ := strconv.Atoi(combine)

			if(sSize >= 13 && sSize <= 24){ // Validation whether the size is more than or equal to 13 AND less than or equal to 24 then add it into the main array

				// book := new(ListBook)
				// book.Category = sCategory
				// book.Title = sTitle
				// book.Size = sSize

				tempArray = append(tempArray, ListBook{
					Category: sCategory,
					Title: sTitle,
					Size: sSize,
				})

			}

		}


	}


	// sort.Sort(doMultiSort(tempArray)) // Sort the array by it's category and size

	// Iterate each splitted array and append it to an array based their category
	for i := range(tempArray) {

		// c6, c7, c0, c9, c4, c8, c1, c2, c5, c3
		itemData := tempArray[i]
		if(itemData.Category == 0){

			arrC0 = append(arrC0, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 1){

			arrC1 = append(arrC1, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 2){

			arrC2 = append(arrC2, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 3){

			arrC3 = append(arrC3, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 4){

			arrC4 = append(arrC4, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 5){

			arrC5 = append(arrC5, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 6){

			arrC6 = append(arrC6, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 7){

			arrC7 = append(arrC7, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 8){

			arrC8 = append(arrC8, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}else if(itemData.Category == 9){

			arrC9 = append(arrC9, ListBook{
				Category: itemData.Category,
				Title: itemData.Title,
				Size: itemData.Size,
			})

		}


		// fmt.Println("Arr:", itemData)
	}

	// End of iteration

	// Iterate each item of Category and append it to an string
	a := []int{6, 7, 0, 9, 4, 8, 1, 2, 5, 3}
	var datapass []ListBook
	for _, s := range a {

		if s == 0 {
			datapass = arrC0
		}else if s == 1 {
			datapass = arrC1
		}else if s == 2 {
			datapass = arrC2
		}else if s == 3 {
			datapass = arrC3
		}else if s == 4 {
			datapass = arrC4
		}else if s == 5 {
			datapass = arrC5
		}else if s == 6 {
			datapass = arrC6
		}else if s == 7 {
			datapass = arrC7
		}else if s == 8 {
			datapass = arrC8
		}else if s == 9 {
			datapass = arrC9
		}



		// sort.Sort(doMultiSort(datapass))


		if tempString == "" {

			process := processBookPerCategory(s, datapass)

			if process == "" {}else{
				tempString = process
			}

		}else{

			process := processBookPerCategory(s, datapass)

			if process == "" {}else{
				tempString = tempString + " " + process
			}

		}
	}
	// End of iteration

	return string(tempString)

}
// End of phase 1

// Second phase, will return the arranged data as string
func processBookPerCategory(category int, data []ListBook) string {

	sort.Slice(data, func(i, j int) bool {

		return data[i].Size > data[j].Size
	})

	// fmt.Println(data)

	returnString := ""

	if len(data) > 0 {

		// countedSimilarData := 1
		// nextDataTitle := ""
		
		sliceKeys := make(map[string]string) // For checking only
		groupedSlices := make(map[string][]ListBook) // Array group for each item with the same Title

		for i := range data {

			if _, ok := sliceKeys[data[i].Title]; ok { // If Title already in array SliceKeys, then add the item to it

				generate := groupedSlices[data[i].Title]
				generate = append(generate, data[i])
				groupedSlices[data[i].Title] = generate

				sort.Sort(doMultiSort(groupedSlices[data[i].Title]))

				// sort.Slice(groupedSlices[data[i].Title], func(i, j int) bool {

					// 	return groupedSlices[data[i].Title][i].Size > groupedSlices[data[i].Title][j].Size
					// }) // Sort the groupped Title by their size from the largest to the smalest

				}else{ // If Title not in array SliceKeys, then create one and add the item to it

					sliceKeys[data[i].Title] = data[i].Title
					generate := []ListBook{}
					generate = append(generate, data[i])
					groupedSlices[data[i].Title] = generate

					sort.Sort(doMultiSort(groupedSlices[data[i].Title]))

					// sort.Slice(groupedSlices[data[i].Title], func(i, j int) bool {

						// 	return groupedSlices[data[i].Title][i].Size > groupedSlices[data[i].Title][j].Size
						// })

					}
					// fmt.Println("===")
					// fmt.Println(groupedSlices[data[i].Title])
					// fmt.Println("===")
				}

				// fmt.Println(sliceKeys)
				countedSimilarData := 1
				nextDataTitle := ""

				// fmt.Println(groupedSlices)
				for i := range groupedSlices { // Iterate the main parent of passed array

					// fmt.Println("vvvvvvvvvvvv")
					// fmt.Println(i)
					// fmt.Println("^^^^^^^^^^^^")

					num := 0
					for ii := range groupedSlices[i] { // Iterate it's subarray (Grouped item by Title) to create an single string


						// fmt.Println(groupedSlices[i][ii])

						// c6, c7, c0, c9, c4, c8, c1, c2, c5, c3
						itemData := groupedSlices[i][ii]
						tempCateg := strconv.Itoa(itemData.Category)
						tempSize := strconv.Itoa(itemData.Size)
						tempArrange := tempCateg + itemData.Title + tempSize
						// fmt.Println(len(groupedSlices[i]))
						if (num+1) >= len(groupedSlices[i]) { 
							if countedSimilarData <= 2 {
								if returnString == ""{
									returnString = tempArrange
								}else{
									returnString = returnString + " " + tempArrange
								}
								countedSimilarData = 1
							}else{
								countedSimilarData = 1
							}
							// fmt.Println("LAST")
						}else{

							itemDataNext := groupedSlices[i][ii+1]
							tempCategNext := strconv.Itoa(itemDataNext.Category)
							// tempSizeNext := strconv.Itoa(itemDataNext.Size)
							nextDataTitle = tempCategNext + itemDataNext.Title
							tempArrangeRN := tempCateg + itemData.Title

							// fmt.Println(tempArrangeRN)
							// fmt.Println(nextDataTitle)


							if nextDataTitle == tempArrangeRN {
								if countedSimilarData <= 2 {
									if returnString == ""{
										returnString = tempArrange
									}else{
										returnString = returnString + " " + tempArrange
									}
									countedSimilarData++
									// fmt.Println("DUP")
								}else{
									countedSimilarData = 1
									// fmt.Println("NON DUP")
								}
							}else{

								if countedSimilarData <= 2 {
									if returnString == ""{
										returnString = tempArrange
									}else{
										returnString = returnString + " " + tempArrange
									}
									nextDataTitle = itemDataNext.Title

									if nextDataTitle == tempArrangeRN {
										countedSimilarData++
									}else{
										countedSimilarData = 1
									}
									// fmt.Println("NON DUP 2 POS")
								}else{
									// fmt.Println("NON DUP 2 NET")
									countedSimilarData = 1
								}


							}

							// fmt.Println(countedSimilarData)


						}

						num++

					}

				}

			}else{
				returnString = ""
			}

			return returnString
		}
		// End of phase 2

		// Function for checking whether string is empty or not
		func checkStringIfEmpty(s string) string {

			if len(s) > 0 {
				return " " + s
			}else{
				return ""
			}

		}

		// Function for converting bytes to string
		func bytesToString(data byte) string {
			return string(byte(data))
		}
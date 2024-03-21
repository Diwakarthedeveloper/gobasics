package main

import "fmt"

func main() {

	//var publisher writer artist title string
	publisher, writer, artist, title := "DizzyBooks Publishing Inc.", "Tracey Hatchet", "Jewel Tampson", "Mr. GoToSleep"
	//var year pageNumber int
	year, pageNumber := "1997", "14"
	//var grade float32
	grade := "6.5"
	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in the year", year, " with total number of pages = ", pageNumber, "which has a grade of", grade)

	////////////////////////////////
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	title = "Epic Vol. 1"
	//writer, artist, title := "Ryan N. Shawn", "Phoebe Paperclips", "Epic Vol. 1" // publisher will be same
	//var year pageNumber int
	//year, pageNumber := "2013", "160"
	year = "2013"
	pageNumber = "160"
	//var grade float32
	grade = "9.0"

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in the year", year, " with total number of pages = ", pageNumber, "which has a grade of", grade)

}

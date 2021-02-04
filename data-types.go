// Data Types

package main

import ( 
	"fmt"
	"strings"
)

func main() {

	//Numbers
	x := 10
	y := 20

	mean := (x + y) / 2.0

	fmt.Println("----------")


	fmt.Println("Mean of x=%v ", x)
	fmt.Println("and of y=%v ", y)
	fmt.Println("is mean=%v \n", mean)

	//Strings

	bookName := "Chanakya's Chant"
	bookAuthor := "Ashwin Sanghi"

	book := "I am reading " + bookName + " by " + bookAuthor + "."

	fmt.Println("----------")

	fmt.Println(book)
	fmt.Println("Name of book is : ", bookName)
	fmt.Println("Author of book is : ", bookAuthor)

	//Slices

	nums := []int{1,2,3,4,5,6}

	fmt.Println("----------")


	fmt.Println("The length of slice is : ", len(nums))
	fmt.Println("The value of slice is : ", nums)

	nums = append(nums, 7)

	fmt.Println("New length of slice is : ", len(nums))


	//Maps

	text := "If you do so, they will become part of your life, and you will gain strong conviction in faith. The teachings that you deeply engrave in your heart will definitely become the foundation for victory in life and enable you to transform your karma"

	words := strings.Fields(text)
	counts := map[string]int{} //create empty map
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println("----------")
	fmt.Println(counts)
}
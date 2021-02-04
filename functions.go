// Data Types

package main

import ( 
	"fmt"
	"math"
)

func getSqrt(num float64) {
	if(num < 0){
		fmt.Printf("ERROR: Number (%f) is a negative number \n", num)
	} else{
		sqrt := math.Sqrt(num)
		fmt.Printf("Square root of %v is %v \n", num, sqrt)
	}
}

func main() {

	fmt.Printf("---------- \n\n")

	getSqrt(-10.0)
	getSqrt(100)

	fmt.Println("\n\n----------")


}
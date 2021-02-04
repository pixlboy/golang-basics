// Channels help go routines send and receive values, below is an example of a buffered channel

package main

import ( 
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n ===== CHANNELS AND GO ROUTINES ====== \n")

	// make is a built in function to create channels
	ch := make(chan int)
	
	//Create a goroutine, then push to channel to send multiple values, close the channel
	go func(){
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}

	fmt.Println("\n ===== CHANNELS AND GO ROUTINES ====== \n")
	
}
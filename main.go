package main


//No-Join
import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
		fmt.Println("Done Waiting for the Main exits")
	}()

	go Work()
	time.Sleep(100 * time.Millisecond)

}

func Work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing my Work to be done")
}

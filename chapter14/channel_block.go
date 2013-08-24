package main
import "fmt"
import "time"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int) 		
	go pump(ch1)		 // pump hangs
	for {		 
		fmt.Println(<-ch1)	
	}
	time.Sleep(1e10)
}
func pump(ch chan int) {
		
	for i:= 2; ;i++ {
		ch <- i
	}
}

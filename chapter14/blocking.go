package main
import (
"time"
"fmt"
)
func f1(in chan int) {
		
fmt.Println(<-in)
}
func main() {
		buf:=2
		 out := make(chan int, buf)
		 out <- 2
		 go f1(out)
	time.Sleep(2e9)
}

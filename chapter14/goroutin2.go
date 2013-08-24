package main
import (
		 "fmt"
//		 "time"
)
func main() {
		
ch := make(chan string)
gch := make(chan int)
var  c ,cc  int = 0, 0
		 go sendData(ch, gch)
		 go getData(ch, gch)
	 	 for {
	             cc= <-gch
                     c=c + cc
                     if cc >= 2 {
			print("finish main\n")
		     	break
		     }
		}
}
func sendData(ch chan string, gch chan int) {
		 ch <- "Washington"
		 ch <- "Tripoli"
		 ch <- "London"
		 ch <- "Beijing"
		 ch <- "Tokio"
		ch <- "end"
			gch <- 1
}
func getData(ch chan string, gch chan int) {
		 var input string
		 for {
			 input = <-ch
			 fmt.Println("%s ", input)
		         if input == "end" {
				break			
			}
		  }
		  gch <- 2
}

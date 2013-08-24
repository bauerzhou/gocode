package main
import "fmt"
import "time"

func main(){
	ch := make(chan string, 1024)
	go sendData(ch)

	time.Sleep(4e9)

	getData(ch)
}

func sendData(ch chan string){
	i:=1
	ch <- "wp"
	ch <- "hello zjqsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"
	for {
		ch <- "test "
		i = i+1
		print(i)
		print(" ")
	}	
	ch <- "\n"
	close(ch)
}

func getData(ch chan string){
	for {
		input, open:= <-ch
		if !open {
			break		
		}
		fmt.Printf("%s ", input)
	}
}

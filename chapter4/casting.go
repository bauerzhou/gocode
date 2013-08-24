package main

import (
 "fmt"
 "time"
 "math/rand" 
 "strings"
)
func main(){
	var n int16 = 34
	var m int32 

	m = int32(n)

 	fmt.Printf("32 bit int is: %d\n",m)
	fmt.Printf("16 bit int is: %d\n", n)

		
	for i := 0; i < 10; i++ {
		 a := rand.Int()
		 fmt.Printf("%d \n", a)
		
	}
	
	for i := 0; i < 5; i++ {
		 r := rand.Intn(8)
		 fmt.Printf("%d \n", r)
	}
	
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f\n", 100*rand.Float32())
 	}
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H’s in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))
	fmt.Printf("Number of double g’s in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}

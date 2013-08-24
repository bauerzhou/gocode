package main

import "fmt"
import "./trans"
import "runtime"

var tp = 2 * trans.Pi
func init(){
    print("main init\n")
}
func main(){
  fmt.Printf("2*Pi = %g\n",tp)
  print("runtime ", runtime.GOOS,"\n")
  str := "hello,i am bauerzhou , good test,good cs , have a lotof fun"
  for i, c := range str {
	fmt.Printf("%d, %c ", i, c)
  }
  fmt.Printf("\n")
}

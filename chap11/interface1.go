package main
import "fmt"

type Animal interface {
	Eat()
	Walk() 
	HaveSex() 
}

type Intelligent interface {
	Speaking()
	Thinking()
}

type Nothing interface {
	GoDie()
	Killing()
}
type Person struct {
	 say string
}
func (p *Person) Eat(){
	print("Eating ", p.say)
}
func (p *Person) Walk(){
	print("walking ", p.say)
}
func (p *Person) HaveSex(){
	print("sexing ", p.say)
}
func (p *Person) Speaking(){
	print("speaking ", p.say)
}
func (p *Person) Thinking(){
	print("Thinking ", p.say)
}
func (p *Person) GoDie(){
	print("wil die ", p.say)
}
func (p *Person) Killing(){
	print("kill each other ", p.say)
}
func main(){
	p := new(Person)
	p.say = "human can do it\n"
	animal := Animal(p)
   	fmt.Println("animal can and: ");  animal.Eat(); animal.Walk(); animal.HaveSex()

	intl := Intelligent(p)
	fmt.Println("intelligent can: ")
	intl.Speaking(); intl.Thinking();

	no := Nothing(p)
	fmt.Println("nothing ...:"); no.GoDie(); no.Killing();
}

package main
import "fmt"
func getHello(name string)string{
	return fmt.Sprintf("Hello,%s",name)
}
func main(){
	s:=getHello("Bob")
	fmt.Println(s)
}
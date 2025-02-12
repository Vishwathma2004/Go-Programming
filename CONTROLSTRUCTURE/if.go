package main
import "fmt"
func main(){
	a:=2000
	if(a<0) {
		fmt.Println("The value is negative")
	} else if(a<10) {
		fmt.Println("The value is less than 10")
	} else {
		fmt.Println("The value is greater than 10")
	}
}
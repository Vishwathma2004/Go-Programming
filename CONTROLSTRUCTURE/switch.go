package main
import "fmt"

func main(){
	a:=10
	switch a==10 {
	case true:
		fmt.Println("Yes!")
	case false:
		fmt.Println("No!")
	}
}
package main
import "fmt"
func main(){
	var x interface{}
	// x=100
	// fmt.Println(x)
	x = "Hello,World"
	// s,ok :=x.(int)
	// if !ok{
	// 	panic("not a integer")
	// }
	// fmt.Println(s)
	switch x.(type){
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	}
}
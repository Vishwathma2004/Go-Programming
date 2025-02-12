//Slices are variable length arrays for storing a single data type


package main
import "fmt"
func main(){
	s:=[]int{1,2,3,4,5}
	fmt.Println(s)
	s = append(s,6)
	fmt.Println(s)


	//we can also create a slice using "make" keyword
	a:=make([]int,0)
	fmt.Println(a)
	b:=make([] int,5)
	b[0]=1
	fmt.Println(b)

	//another method
	c:=[4]int{1,2,3,4}
	fmt.Println(c)
}
package main
import "fmt"
type Person struct{
	First string
	Last string
	Age int
	Phone Phone
}
type Phone struct{
	Areacode string
	Prefix string
	Suffix string
}
func main(){
	p:=Person{
		First:"Allen",
		Last:"Bob",
		Age: 60,
		Phone:Phone{
			Areacode:"123",
			Prefix:"456",
			Suffix:"789",	
		},
	}
	// q:=Person{"Who","are",50}
	fmt.Println(p)

}
//Maps are key value pairs so that single values can be looked up quickly

package main
import "fmt"
func main(){
	m:=map[string]string{}
	m["first"]="allen"
	m["last"] = "bob"
	fmt.Println(m)

	//Second Method - we also supply the values we want

	n:=map[string]string{
		"first":"cat",
		"last":"dog",
	}
	fmt.Println(n)


	//Third Method - Using "Slice Keyword"

	a:=make(map[string]string)
	a["first"] = "egg"
	a["last"] = "fish"
	fmt.Println(a)
	
	fmt.Println(a["first"])
}
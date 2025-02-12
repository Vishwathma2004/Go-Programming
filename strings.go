package main
import(
	"fmt"
	"strings"
)
func main(){
	s:="I'm a string"
	fmt.Printf("Ends with a String?%v\n",strings.HasSuffix(s,"string"))
}
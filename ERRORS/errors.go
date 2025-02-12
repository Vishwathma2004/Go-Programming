package main
import (
	"fmt"
	"strconv"
	"os"
)
func main(){
	var sum int
	for _, a:=range os.Args[1:]{
		i,err:=strconv.Atoi(a)
		if err!=nil{
			panic(fmt.Sprintf("Invalid value:%v",err))
		}
		sum+=i
	}
	fmt.Printf("Sum =%v",sum)
}
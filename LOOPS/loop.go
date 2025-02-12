package main
import (
	"fmt"
	"time"
)
func main(){
	fmt.Println("Tick")
	time.sleep(1*time.Second)
	fmt.Println("Tock")
	time.sleep(1*time.Second)
}
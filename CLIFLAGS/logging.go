//go provide a standard library for logging which will allow you to log at different levels and to also log fatal errors
package main
import "log"
func main(){
	log.Println("Hello World")
	// log.Panic("oops!")
	log.Fatalln("oops there is an error")
}
//Types such as structure can have method sets
//This means that the type has function that can be called on an instance of the type
//This is similar to how classes work in other languages
package main
import "fmt"
func sayHello(name string){
    fmt.Printf("Hello, %s\n",name)
}
func main(){
    sayHello("Bob")
}
//Interface provide an abstraction so that different types that implement 
//the same method can be treated the same way

package main
import "fmt"
type Employee struct{
	BasePerson
	Salary inn
	Boss *Manager
}
type Manager struct{
	Employee
}
func main(){
	m:= &Manager{
		Employee{
			BasePerson:BasePerson{
				First:"John",
				Last:"Doe",
			},
			Salary:100000,
		},
	}
	e:=&Employee{
		BasePerson:BasePerson{
			First:"Jane",
			Last:"Doe",
		}
	}
}
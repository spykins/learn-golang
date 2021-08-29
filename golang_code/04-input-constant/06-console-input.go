package main
import "fmt"

func main() {
  fmt.Println("Enter Your First Name: ")
  var firstName string 
  fmt.Scanln(&firstName)

  fmt.Println("Enter Middle Name: ")
  var middleName string
  fmt.Scanln(&middleName)
  
  fmt.Println("The firstName = " + firstName + "and the middle name = " + middleName)

}

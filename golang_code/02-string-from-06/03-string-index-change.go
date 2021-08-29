// Error ....... string are immutable
package main 
import "fmt"

func main() {
   mystr := "Welcome to Go Lang World!!"
   fmt.Println("String: ", mystr)
  
   // mystr[1]  = 'G' // Strings are immutable
   fmt.Println("String: ", mystr) 

}

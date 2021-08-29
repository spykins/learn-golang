package main 
import (
   "fmt"
   "unicode/utf8"
)


func main() {
  mystr := "Welcome to Go Language World"

  length1 := len(mystr)

  length2 := utf8.RuneCountInString(mystr) // another way  
  fmt.Println("string: ", mystr)
  fmt.Println("Length 1: ", length1)
  fmt.Println("Length2: ", length2)
}

package main
import "fmt"

 const (
   PRODUCT string = "Website"
   QUANTITY = 1
   PRICE = 42.42
   ONLINE = true
 )

func main() {
  // all varialbes names must be in camel case
 // and constants must be in Capital letters.

  fmt.Println(QUANTITY)
  fmt.Println(PRODUCT)
  fmt.Println(PRICE)
  fmt.Println(ONLINE) 
}

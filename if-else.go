package main
import "fmt"

func main(){

  if 7%2 == 0 {
    fmt.Println("The number is even")
  } else {
    fmt.Println("The number is odd")
  }

  if 8%2 == 0 {
    fmt.Println("The number is divisible by 2")
  }

  if 8%2 == 0 && 7%2 == 0 {
    fmt.Println("Both are even")
  } else {
    fmt.Println("One of them is odd")
  }

  if num:= 9 ; num < 0 {
    fmt.Println( num , "is negative")
  } else if num < 10 {
    fmt.Println(num , "is positive 1 digit number")
  } else {
    fmt.Println(num , "is positive multi digit number")
  }
}

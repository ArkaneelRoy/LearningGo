package main
import (
  "fmt"
  "time"
       )

func main() {
  i := 2
  fmt.Println("Print" , i , "as:" )
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday , time.Sunday :
    fmt.Println("Weekend Baby!")
  default:  
    fmt.Println("Its a boring Weekday (TT)")
  }

  t := time.Now()
  switch  {
  case t.Hour() < 12:
    fmt.Println("Its before noon")
  default:
    fmt.Println("Its after noon")
  }
  
  whatami := func(i any) {
    switch t := i.(type) {
    case bool :
      fmt.Println("I am boolean")
    case int :
      fmt.Println("I am integer")
    default :
      fmt.Printf("I dont know my type %T\n" , t )
    }
  }
  whatami(true)
  whatami(1)
  whatami("hey")
}

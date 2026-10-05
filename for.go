package main 
import "fmt"

func main(){
   i := 1
  for i <= 4 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 0 ; j < 4 ; j++ {
    fmt.Println(j)
  }

  for i := range 5 {
    fmt.Println("Range:", i)
  }

  for {
    fmt.Println("Infinite Loop")
    break
  }

  for n := range 6 {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}

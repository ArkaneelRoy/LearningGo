package main 
import (
  "fmt" 
  "math"
)

const s string  = "A constant string value"
func main(){
  fmt.Println(s)

  const a = 50000000000
  const b = 5e20/a

  fmt.Println(b)
  fmt.Println(int64(b))
  fmt.Println(math.Sin(a))
}

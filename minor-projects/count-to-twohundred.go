package main
import  "fmt"
func main()  {
  
  for j := 0 ; j >= 0 ; j++ {
    if j <= 200 {
      fmt.Println(j)
    } else {
      fmt.Println("End")
      break
    }
  }
}

package main 
import "fmt"
func main() {

 /** fmt.Println("BENCHMARK START..")
  var value int = 0
  for value >= 0 {
    if value <= 10000000000 {
      value = value + 1
      //fmt.Println(value)
    } else{
      fmt.Println("BENCHMARK END")
      break
    }
  } **/

  //idiomatic syntax 
  
fmt.Println("BENCHMARK START...")
  for value := 0 ; value <= 1_000_000_000 ; value++{
    //counts
  }
fmt.Println("BENCHMARK END...")  
}

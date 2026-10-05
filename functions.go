package main
import "fmt"

func add(a int , b int) int {
  return a + b
}

func greet(){
  fmt.Println("So long suckers!")
}

func PrintSquare(n int){
  fmt.Println("Square is :" , n*n)
}

func divide(a int , b int) (int , int){
  quotient:= a/b
  remainder:=a%b
  return quotient , remainder
}

func main (){
  result1:=add(10,20)
  result2:=add(32,8)

  fmt.Println(result1 , "and" , result2)
  greet() //Output of greet function
  PrintSquare(25) //Output of PrintSquare function

  q , r := divide(12,5)
  fmt.Println("Quotient is:" , q , "and Remainder is :" , r)
}

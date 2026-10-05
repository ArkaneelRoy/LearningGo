package main 
import "fmt"

func main(){
  var age int
  fmt.Print("Enter your number :")
  fmt.Scanln(&age)
  fmt.Println("Your age is ", age)

  var num1 , num2 int
  fmt.Print("Enter 2 numbers to add with space :")
  fmt.Scanln(&num1 , &num2)
  fmt.Println("The sum is :" , num1+num2)
}

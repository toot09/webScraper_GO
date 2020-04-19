package main

import ("fmt"
        //"strings"
        )

// Constructure 선언
type person struct{
  name    string
  age     int
  favFood []string
}
func main() {
  favFood := []string{"pizza","chicken"}
  // Constructure 사용
  jaewon := person{name:"jawon", age:32, favFood:favFood}
  fmt.Println(jaewon.name)
}
// compile이 필요한 프로젝트라면 시작 패키지는 main 이름을 가짐
// package 선언
package main

import "fmt"

// 제일먼저 실행되는 function이 main function
func main() {
  // 대문자로 function을 만들면 다른 프로젝트에서 export 가능
  //fmt.Print("안녕하세요")
  
  //constance(상수) : constance는 변경 불가능
  const name_c string = "jaewon"
  fmt.Println(name_c)

  //variable(변수) : variable은 변경 가능 
  //축약형 :=로 선언가능 ,type 자동매칭 가능 (function 안에서만 가능)
  //var name_v string = "jaewon"
  name_v := "ayoung"
  fmt.Println(name_v)
  chk := false
  fmt.Println(chk)
}
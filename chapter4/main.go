package main

import "fmt"

const c string = "C is a constant"

func main () {
  var x string = "Hello World"
  fmt.Println(x);
  var y string
  y = "Welcome"
  y +=" to go!!!"
  fmt.Println(y)
  z := "Go rocks!!"
  fmt.Println(z)
  fmt.Println(c)
}

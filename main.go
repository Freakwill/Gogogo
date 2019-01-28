
package main

import "C"

// go build -buildmode=c-shared -o path/test.so path/main.go

import (
   "fmt"
   "C"
)

/* 定义结构体 */
type Circle struct {
  Radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.Radius 即为 Circle 类型对象中的属性
  return 3 * c.Radius * c.Radius
}

func (c Circle) setRadius(r float64) {
  //c.Radius 即为 Circle 类型对象中的属性
  c.Radius = r
}

//export test
func test(r float64) {
  var c1 = Circle{Radius: r}
  fmt.Println("r = ", r)
  fmt.Println("r = ", c1.Radius)
  fmt.Println("S = ", c1.getArea())
}

func main() {
    //test(3)
}
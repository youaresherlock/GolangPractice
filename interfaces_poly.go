/*
* @Author: Clarence
* @Date:   2019-07-13 18:46:59
* @Last Modified by:   Clarence
* @Last Modified time: 2019-07-13 19:39:46
*/
/*
接口是一种契约，实现类型必须满足它,它描述了类型的行为，规定类型可以做什么.接口彻底将类型
能做什么，自己如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为,这就是多态的本质
*/
package main

import "fmt"

type Shaper interface{
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}

	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

}


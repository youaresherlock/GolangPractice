/*
* @Author: Clarence
* @Date:   2019-09-12 09:53:16
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-12 09:59:52
*/
/*
defer声明会先计算确定参数的值, defer推迟执行的仅仅是函数体 
*/
package main 

import (
	"fmt"
	"time"
)

func main() {
	i, p := a()
	fmt.Println("return: ", i, p, time.Now())
}

func a() (i int, p *int) {
	// i++
	defer func(i int) {
		fmt.Println("defer3: ", i, &i, time.Now())
	}(i)

	defer  fmt.Println("defer2: ", i, &i, time.Now()) 

	defer func() {
		fmt.Println("defer1: ", i, &i, time.Now())
	}()

	i++

	func() {
		fmt.Println("print1: ", i, &i, time.Now())
	}() 

	fmt.Println("print2: ", i, &i, time.Now())

	return i, &i
}

// print1:  1 0xc000066058 2019-09-12 09:57:18.4133357 +0800 CST m=+0.012962501
// print2:  1 0xc000066058 2019-09-12 09:57:18.6517226 +0800 CST m=+0.251349401
// defer1:  1 0xc000066058 2019-09-12 09:57:18.6537174 +0800 CST m=+0.253344201
// defer2:  0 0xc000066058 2019-09-12 09:57:18.4133357 +0800 CST m=+0.012962501
// defer3:  0 0xc000066818 2019-09-12 09:57:18.6557122 +0800 CST m=+0.255339001
// return:  1 0xc000066058 2019-09-12 09:57:18.6567095 +0800 CST m=+0.256336301
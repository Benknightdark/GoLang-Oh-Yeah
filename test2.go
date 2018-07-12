// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

const (
	i      = 100
	pi     = 444
	prefix = "go_"
)

const (
	x = iota
	y = iota
	z = iota
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
func p(s interface{}) {
	fmt.Println(s)
}

func main() {
	fmt.Println(i)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(z)
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(doubleArray[0])
	var a = [10]byte{'a', 'b'}
	p(string(a[0]))
	m := make(map[string]string)
	m["hello"] = "bbb"
	p(m["hello"])
	delete(m, "hello")
	p(m)
	m1 := m
	m1["hello"] = "ccc"
	p(m)
	p(m1)

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	p(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	p(sum2)

	for k, v := range m {
		p(k)
		p(v)
	}

	for _, v := range m {

		p(v)
	}
	x := 1
	y := 2
	z := 3
	p(max(x, y))
	p(max(z, y))
	sumxy, meanxy := SumAndProduct(x, y)
	p(sumxy)
	p(meanxy)

}

// You can edit this code!
// Click here and start typing.
package main

import (
	. "fmt"
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
	Println(s)
}
func add1(a *int) int {
	*a = *a + 1
	return *a
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true

}
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false

}
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func throwPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

type person struct {
	name string
	age  int
}

func main() {
	var ap person
	ap.name = "aa"
	ap.age = 22
	p(ap)
	// user := os.Getenv("USERaaa")
	// // throwPanic()
	// if user == "" {
	// 	panic("no user env")
	// }
	p("-----------將FUNCTION包成參數呼叫---------------")
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	odd := filter(slice, isOdd)
	p(odd)
	even := filter(slice, isEven)
	p(even)
	p("---------------------------------")
	Println(i)
	Println(x)
	Println(y)

	Println(z)
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	Println(doubleArray[0])
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
	x1 := add1(&x)
	p(x)
	p(x1)
	for i := 0; i < 5; i++ {
		defer Printf("%d ", i)
	}

}

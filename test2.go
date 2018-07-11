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

}
func p(s interface{}) {
	fmt.Println(s)
}

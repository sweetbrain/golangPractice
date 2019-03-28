package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := [3]string{
		"Yamada Tarou",
		"Sato Hanako",
		"Suzuki Souta",
	}

	fmt.Println(a)
	fmt.Printf("%d年%d月%d日\n", 2019, 04, 10)
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 4, "Golang", [...]int{1,2,3,4})

	var x,y,z int
	x, y, z = 1,23,4
	fmt.Print(x,y,z)

	var (
		n = 1
		s = "string"
		b = true
	)

	//n := 1
	//s := "string"
	//b := true

	if b == true {
		fmt.Print(n, s)
	}

	name := "Taro" + " " + "¥amada"
	fmt.Printf("%v\n", name)

	fmt.Print(plus(1, 0))

	fmt.Print(div(20, 3))

	f := returnFunc()
	f()

	l := later()

	fmt.Println(l("Golang"))
	fmt.Println(l("is"))
	fmt.Println(l("awesome!"))

	xx, yy := 2, 5
	if n := xx * yy; n%2 == 0 {
		fmt.Printf("n(%d) is even\n", n)
	} else {
		fmt.Printf("n(%d) is odd\n", n)
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	go fmt.Println("Yeah!")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())

	ss := make([]int, 10)
	fmt.Println(ss)

	var aa [10]int
	fmt.Println(aa)


}

func plus(x, y int) int {
	return x + y
}

func div(a, b int) (int ,int) {
	q := a / b
	r := a % b

	return q,r
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func init() {
	fmt.Println("init()")
}



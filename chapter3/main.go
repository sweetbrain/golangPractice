package main

import "fmt"

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

}

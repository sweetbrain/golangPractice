package main

import (
	"fmt"
)

func main() {

	//slice := make([]float64, 3)
	//fmt.Println(slice)
	//slice[0] = 3.14
	//fmt.Println(slice)
	//slice[1] = 6.28
	//fmt.Println(slice)
	//fmt.Println(slice[0])
	//
	//s1 := make([]int, 8)
	//fmt.Println(len(s1))
	//
	//s2 := make([]int, 5)
	//fmt.Println(len(s2))
	//fmt.Println(cap(s2))
	//
	//s3 := make([]int, 5, 10)
	//fmt.Println(len(s3))
	//fmt.Println(cap(s3))
	//
	//s0 := []int{1, 2, 3}
	//s4 := []int{4, 5, 6}
	//s5 := append(s0, s4...)
	//fmt.Println(s5)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s7 := a[2:4]
	fmt.Println(len(s7))
	fmt.Println(cap(s7))
	fmt.Println(s7)

	s8 := a[2:4:4]
	fmt.Println(len(s8))
	fmt.Println(cap(s8))
	fmt.Println(s8)

	s9 := a[2:4:6]
	fmt.Println(len(s9))
	fmt.Println(cap(s9))
	fmt.Println(s9)

	s := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s {
		fmt.Printf("[%d] -> %s\n", i, v)
		s = append(s, "Melon")
	}

	fmt.Println(s)
	fmt.Println(sum(1, 4, 6, 7, 5))

	s1 := []int{1, 2, 3}
	fmt.Println(sum(s1...))

	ar := []int{1, 2, 3}
	pow(ar)
	fmt.Println(ar)

	var (
		k [3]int
		l []int
	)
	fmt.Println(k)
	fmt.Println(l == nil)

	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	ms := make(map[string]string)

	ms["Yamada"] = "Taro"
	ms["Sato"] = "Yuuki"
	ms["Yamada"] = "Jiro"

	fmt.Println(ms)

	mk := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(mk)

	mj := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}

	fmt.Println(mj)

	mh := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}
	fmt.Println(mh[1][3.14])

	if _, ok := m[81]; ok {
		fmt.Print("通過\n")
	}

	mb := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}

	sa := mb[1]
	if sa != nil {
		fmt.Println("This is not nil")
	}

	fruits := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	for k, v := range fruits {
		fmt.Printf("%d => %s\n", k, v)
	}

	md := map[string]int{
		"Apple":  88,
		"Banana": 107,
		"Cherry": 46,
	}
	md["Grape"] = 66
	md["Lemon"] = 16
	md["Orange"] = 44
	md["pineapple"] = 73

	for k, v := range md {
		fmt.Printf("%s => %d\n", k, v)
	}

	//ch := make(chan int, 10)
	//
	//ch <- 5
	//i := <-ch
	//fmt.Println(i)

	//ch := make(chan int)
	//
	//go receiver(ch)
	//
	//i := 0
	//for i < 10 {
	//	ch <- i
	//	i++
	//}

	//ch := make(chan string, 3)
	//
	//ch <- "Apple"
	//fmt.Println(len(ch))
	//ch <- "Banana"
	//ch <- "Cherry"
	//fmt.Println(len(ch))

	//ch := make(chan int, 20)
	//
	//go receive("1st goroutine", ch)
	//go receive("2nd goroutine", ch)
	//go receive("3nd goroutine", ch)
	//go receive("4th goroutine", ch)
	//go receive("5th goroutine", ch)
	//
	//i := 0
	//for i < 100 {
	//	ch <- i
	//	i++
	//}
	//close(ch)
	//
	//time.Sleep(3 * time.Second)

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//ch1 <- 1
	//ch2 <- 2

	//select {
	//case <-ch1:
	//	fmt.Println("ch1から受信")
	//case <-ch2:
	//	fmt.Println("ch2から受信")
	//case ch3 <- 3:
	//	fmt.Println("ch3へ送信")
	//default:
	//	fmt.Println("ここへは到達しない")
	//}
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func sum(s ...int) int {
	n := 0

	for _, v := range s {
		n += v
	}

	return n
}

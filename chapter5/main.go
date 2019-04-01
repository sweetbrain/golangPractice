package main

import (
	"fmt"
	"math"
)

//func main() {
//	const totalExecuteNum = 6
//	const maxConcurrency = 3
//	sig := make(chan string, maxConcurrency)
//	res := make(chan string, totalExecuteNum)
//	defer close(sig)
//	defer close(res)
//
//	fmt.Printf("start concurrency execute %s \n", time.Now())
//	for i := 0; i < totalExecuteNum; i++ {
//		go wait6sec(sig, res, fmt.Sprintf("no%d", i))
//	}
//	for {
//		if len(res) >= totalExecuteNum {
//			break
//		}
//	}
//
//	fmt.Printf("end concurrency execute %s \n", time.Now())
//}
//
//func wait6sec(sig chan string, res chan string, name string){
//	sig <- fmt.Sprintf("sig %s", name)
//	time.Sleep(6 * time.Second)
//	fmt.Printf("%s:endwait 6 sec \n", name)
//	res <- fmt.Sprintf("sig %s", name)
//	<-sig
//}

func main() {
	//var i int
	//p := &i
	//fmt.Printf("%T\n", p)
	//pp := &p
	//fmt.Printf("%T\n", pp)

	//var i int
	//p := &i
	//i = 5
	//fmt.Println(*p)
	//*p = 10
	//fmt.Println(i)

	//i := 1
	//inc(&i)
	//inc(&i)
	//inc(&i)
	//fmt.Println(i)

	//p := &[3]int{1, 2, 3}
	//pow(p)
	//fmt.Println(p)

	//a := [3]string{"Apple", "Banana", "Cherry"}
	//p := &a
	//fmt.Println(a[1])
	//fmt.Println(p[1])
	//p[2] = "Grape"
	//fmt.Println(a[2])
	//fmt.Println(p[2])

	//p := &[3]int{1, 2, 3}
	//
	//fmt.Println(len(p))
	//fmt.Println(cap(p))
	//fmt.Println(p[0:2])

	//p := &[3]string{"Apple", "Banana", "Cherry"}
	//
	//for i, v := range p {
	//	fmt.Println(i, v)
	//}

	//i := 5
	//ip := &i
	//fmt.Printf("type=%T, address=%p\n", ip, ip)

	//is := [3]int{1, 2, 3}
	//ip := &is[1]
	//*ip = 0
	//fmt.Println(is)
	//
	//fs := []float64{1.1, 2.2, 3.3}
	//fp := &fs[2]
	//*fp = 3.14
	//fmt.Println(fs)

	//s := ""
	//for _, v := range []string{"A", "B", "C"} {
	//	s += v
	//}
	//fmt.Println(s)

	//s := "Hello, World!"
	//printString(s)

	//type MyInt int
	//
	//var n1 MyInt = 5
	//n2 := MyInt(7)
	//fmt.Println(n1)
	//fmt.Println(n2)

	//type (
	//	IntPair		[2]int
	//	Strings		[]string
	//	AreaMap		map[string][2]float64
	//	IntsChannel	chan [] int
	//)
	//
	//pair := IntPair{1, 2}
	//strs := Strings{"Apple", "Banana", "Cherry"}
	//amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	//ich  := make(IntsChannel)

	//n := Sum(
	//	[]int{1, 2, 3, 4, 5},
	//	func(i int) int {
	//		return i * 2
	//	},
	//)
	//fmt.Println(n)
	//
	//type Person struct {
	//	ID   uint
	//	name string
	//	部署  string
	//}
	//
	//p := Person{ID: 17, name: "山田太郎", 部署: "営業部"}
	//
	//fmt.Println(p)

	//type Feed struct {
	//	Name	string
	//	Amount	uint
	//}
	//
	//type Animal struct {
	//	Name string
	//	Feed
	//}
	//
	//a := Animal{
	//	Name:	"Monkey",
	//	Feed:	Feed{
	//		Name:	"Banana",
	//		Amount:	10,
	//	},
	//}
	//
	//fmt.Println(a.Name)
	//fmt.Println(a.Feed.Name)
	//fmt.Println(a.Amount)
	//
	//a.Amount = 15
	//fmt.Println(a.Amount)

	//type Base struct {
	// 	Id		int
	// 	Owner	string
	//}
	//
	//type A struct {
	//	Base
	//	Name string
	//	Area string
	//}
	//
	//type B struct {
	//	Base
	//	Title	string
	//	Bodies []string
	//}
	//
	//a := A{
	//	Base: Base{17, "Taro"},
	//	Name: "Taro",
	//	Area: "Tokyo",
	//}
	//
	//b :=  B{
	//	Base: Base{81, "Hanako"},
	//	Title: "no title",
	//	Bodies: []string{"A", "B"},
	//}
	//
	//fmt.Println(a.Id)
	//fmt.Println(a.Owner)
	//fmt.Println(b.Id)
	//fmt.Println(b.Owner)

	//s := struct { X, Y int }{ X: 1, Y: 2 }
	//showStruct(s)

	//type Point struct {
	//	X, Y int
	//}
	//
	//p := Point{X: 3, Y: 8}
	//showStruct(p)

	//p := Point{X: 1, Y: 2}
	//swap(&p)
	//
	//fmt.Println(p.X)
	//fmt.Println(p.Y)

	//type Person struct {
	//	Id		int
	//	Name 	string
	//	Area	string
	//}
	//
	//p := new(Person)
	//
	//fmt.Println(p.Id, p.Name, p.Area)

	//p := new(Point)
	//p.X = 1
	//p.Y = 2

	//p := &Point{X: 5, Y: 12}
	//p.Render()

	//p := &IntPoint{X: 0, Y: 0}
	//fmt.Println(p.Distance(&IntPoint{X: 1, Y: 1}))
	//pf := &FloatPoint{X: 0, Y: 0}
	//fmt.Println(pf.Distance(&FloatPoint{X: 1, Y: 1}))

	//ip := IntPair{1, 2}
	//ip.First()
	//ip.Last()
	//
	//fmt.Println(Strings{"A", "B", "C"}.Join(","))

	//fmt.Println(NewUser(1,"Taro"))

	//f := (*Point).ToString
	//fmt.Println(f(&Point{X: 7, Y:11}))

	//p := &Point{X: 2, Y: 3}
	//f := p.ToString
	//fmt.Println(f())

	//p1 := Point{}
	//p1.Set(1,2)
	//
	//fmt.Println(p1.X)
	//fmt.Println(p1.Y)
	//
	//
	//p2 := &Point{}
	//p2.Set(1,2)
	//
	//fmt.Println(p2.X)
	//fmt.Println(p2.Y)

	//ps := make([]Point, 5)
	//for _, p := range ps {
	//	fmt.Println(p.X, p.Y)
	//}

	//ps := Points{}
	//ps = append(ps, &Point{X: 1, Y: 2})
	//ps = append(ps, nil)
	//ps = append(ps, &Point{X: 3, Y: 4})
	//fmt.Println(ps.ToString())

	//m1 := map[User]string {
	//	{Id: 1, Name: "Taro"}: "Tokyo",
	//	{Id: 2, Name: "Jiro"}: "Osaka",
	//}
	//
	//m2 := map[int]User {
	//	1: {Id: 1, Name: "Taro"},
	//	2: {Id: 2, Name: "Jiro"},
	//}
	//
	//fmt.Println(m1)
	//fmt.Println(m2)

	//ms := map[int][]string {
	//	1: {"A", "B", "C"},
	//}
	//
	//mm := map[int]map[int]string {
	//	1: {1: "Apple", 2: "Banana", 3: "Cherry"},
	//}
	//
	//fmt.Println(ms)
	//fmt.Println(mm)

	//u := User{Id: 1, Name: "Taro", Age: 32}
	//
	//t := reflect.TypeOf(u)
	//
	//for i := 0; i < t.NumField(); i++ {
	//	f := t.Field(i)
	//	fmt.Println(f.Name, f.Tag)
	//}

	//u := User{Id: 1, Name: "Taro", Age: 32}
	//bs, _ := json.Marshal(u)
	//fmt.Println(string(bs))

	//err := RaiseError()
	//fmt.Println(err.Error())
	//
	//e, ok := err.(*MyError)
	//if ok {
	//	fmt.Println(e.ErrCode)
	//}

	//vs := []Stringify{
	//	&Person{Name: "Taro", Age: 21},
	//	&Car{Number: "XXX-0123", Model: "PX512"},
	//}
	//
	//for _, v := range vs {
	//	fmt.Println(v.ToString())
	//}

	//Println(&Person{Name: "Hanako", Age: 21})
	//Println(&Car{Number: "XYZ-1234", Model: "RT-37"})

	//t := &T{Id: 10, Name: "Taro"}
	//fmt.Println(t)

	//t := foo.NewI()
	//fmt.Println(t.Method1())

	t := &T{Id: 19}
	ShowId(t)

}

type T struct {
	Id int
}

func (t *T) GetId() int { return t.Id }

func ShowId(id interface{ GetId() int }) {
	fmt.Println(id.GetId())
}

type I0 interface {
	Method1() int
}

type I1 interface {
	Method2() int
}

type I2 interface {
	Method3() int
}

//func (t *T) String() string {
//	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
//}

//type T struct {
//	Id		int
//	Name	string
//}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました。", ErrCode: 1234}
}

//type error interface {
//	Error() string
//}

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d,%d]", p.X, p.Y)
		}
	}
	return str
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

type Point struct {
	X, Y int
}

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

type User struct {
	Id   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  uint   `json:"user_age"`
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Last() int {
	return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

type IntPoint struct{ X, Y int }
type FloatPoint struct{ X, Y float64 }

func (p *IntPoint) Distance(dp *IntPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(x*x + y*y)
}

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (p *Point) Render() {
	fmt.Printf("<%d,%d>\n", p.X, p.Y)
}

func swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func printString(s string) {
	fmt.Println(s)
}

func inc(p *int) {
	*p++
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

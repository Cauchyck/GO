package note

import (
	"fmt"
	"hello_go/util"
	"sync"
)

func SayHelloWorld() {
	fmt.Println("Hello World")
}

func VariablesAndConstant() {
	fmt.Println("\n1 常量")
	var v1 int
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4

	var (
		v5     = 5
		v6 int = 6
		v7 int
	)

	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v, v7=%v\n", v1, v2, v3, v4, v5, v6, v7)
	fmt.Println("\n2 变量")

	const (
		c1 = 8
		c2 = iota
		c3 = iota
		c4
		c5 = 12
		c6
		c7
	)
	fmt.Printf("c1=%v, c2=%v, c3=%v, c4=%v, c5=%v, c6=%v, c7=%v\n", c1, c2, c3, c4, c5, c6, c7)
}


// 指针
func Pointer() {

	var increase = func (n *int) {
		*n++
		fmt.Printf("n = %v\n n's address: %v\n n's value: %v\n", *n, &n, n)
	}
	
	var src = 2022
	var ptr = &src
	increase(ptr)
	fmt.Printf("src = %v\n s's address %v\n  ptr's address %v\n", src, &src, ptr)
}

// fmt格式字符

func FmtVerbs() {
	i := 123
	fmt.Printf("%U \n", i)
	fmt.Printf("%c \n", i)
	fmt.Printf("%q \n", i)

	f := 3.1415926

	fmt.Printf("%f \n", f)
	fmt.Printf("%.2f \n", f)
	fmt.Printf("%5f \n", f)

	fmt.Printf("%b \n", f)
	fmt.Printf("%e \n", f)
	fmt.Printf("%x \n", f)
}

// if...else
func IfElse() {
	var age uint8
	fmt.Println("请输入你的年龄：")

	fmt.Scanln(&age)
	if age < 13 {
		fmt.Print("不要编程！\n不要编程！\n不要编程！\n")
	} else {
		fmt.Println("认命吧。")
	}
}

// switch..case
func SwitchCase() {
	var age uint8
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	switch {
	case age < 10:
		fmt.Print("不要编程！\n不要编程！\n不要编程！\n")
		fallthrough
	case age < 25:
		fmt.Println("认命吧。")
	default:
		fmt.Println("...")

	}

}

// for
func For() {
	i := 1
	fmt.Println("无限循环")
	for {
		fmt.Print(i, "\t")
		i++
		if i == 11 {
			fmt.Println()
			break
		}
	}

	fmt.Println("条件循环")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()
	fmt.Println("标准循环")

	for i := 1; i < 11; i++ {
		fmt.Print(i, "\t")
	}
}

// label and goto
func LabelAndGoto() {
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")

			if i == 9 && j == 4 {
				break outside
			}
		}
		fmt.Println()
	}
}

// 函数
func getRes(n1, n2 int) (sum, diff int){
	sum = n1 + n2
	diff = n1 - n2
	return
}
func Function(){
	sum, diff := getRes(2,3)
	fmt.Println("sum=", sum, ", diff=", diff)

	// fmt.Printf("getRes=%v, Type of getRes=%T\n", getRes, getRes)
}

// defer 延迟执行
// 延迟执行的函数会被压入栈中。 return后安装先进后出的顺序调用
// 延迟执行的函数其参数会立即求值
func deferUtil() func(int) int {
	i := 0
	return func(n int) int{
		fmt.Printf("n = %v\n", n)
		i++
		fmt.Printf("i = %v\n", i)
		return i
	}
}

func Defer() int {
	f := deferUtil()

	defer f(1)
	defer f(2)
	defer f(3)
	defer f(4)

	return f(5)
}

 
func DeferRecover() {
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()

	n := 0
	fmt.Println(3 / n)
}

// 数组
func Array() {
	var a = [...]int{
		1,
		2,
		3,
	}
	for i := 0; i < len(a); i++{
		fmt.Printf("a[%v] = %v\t", i, a[i])
	}
	fmt.Println()
	for i, v := range a {
		fmt.Printf("a[%v] = %v \t", i, v)
	}
	fmt.Println()

	var twoDimensionalArray [3][4] int
	twoDimensionalArray[0][0] = 1
	twoDimensionalArray[1][1] = 2
	twoDimensionalArray[2][2] = 3

	for i, v1 := range twoDimensionalArray{
		for j, v2 := range v1{
			fmt.Printf("b[%v][%v] = %v \t", i, j, v2)
		}
		fmt.Println()
	}
}

// 切片是对数组的引用
// 切片本身并不存储任何数据， 它只是描述了底层数组中的一段
// 索引从0开始
// 切片是引用类型，默认值为nil
// 遍历方式和数组相同
// 切片引用切片仍指向同一个数组
func Slice(){
	array := [5]int{1,2,3,4,5}
	var s1 []int = array[0:3] // 左闭右开
	s1[0] = 0
	fmt.Println("array =", array)
	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array =", array)
	var s3 []int
	fmt.Println("Does s3 == nil?", s3 == nil)
	s3 = make([]int, 3, 5) // make([]Type, len, cap) cap默认与len相同
 	fmt.Printf("len(s3)=%v, cap[s3] = %v", len(s3), cap(s3))

	s4 := []int{1, 2, 4}
	fmt.Println("s4 = ", s4)

	s1 = append(s1, 6, 7, 8) //底层创建了新数组，不再引用原数组
	s1[4] = 0

	fmt.Println("array=", array)
	fmt.Println("s1=", s1)

	s5 := append(s1, s2...)
	fmt.Println("s5 = ", s5)


	s6 := []int{1, 1}
	copy(s6, s5)
	fmt.Println("s6 = ", s6)

	str := "hello 世界"
	fmt.Printf("[]byte(str) = %v\n []bety(str)=%s\n", []byte(str), []byte(str))
	for i, v := range str{
		fmt.Printf("str[%d] = %c \n", i, v)

	}

	key := util.SelectByKey("注册", "登录", "退出")

	fmt.Println("接收到key = ", key)

}

// map
func Map(){
	var m1 map[string]string
	fmt.Println("m1 == nil ?", m1 == nil)

	m1 = make(map[string]string)
	m1["早上"] = "敲代码"
	m1["中午"] = "送外卖"
	m1["晚上"] = "跑滴滴"

	m2 := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}

	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	v, ok := m2["key1"]
	if ok {
		fmt.Println("v = ", v)
	}else{
		fmt.Println("key1 not in m2")
	}

	delete(m1, "晚上")
	fmt.Println("m1 = ", m1)

	for key, value := range m2{
		fmt.Printf("m2[%v] = %v \n", key, value)
	}

}

// 自定义数据类型&类型别名
func TypeDefintionAndTypeAlias(){
	type mesType uint16
	var u1000 uint16=1000
	var textMes mesType = mesType(u1000)
	fmt.Printf("textMes=%v, Type of textMes=%T \n", textMes, textMes)

	type myUint16 = uint16
	var myu16 myUint16 = 1000
	fmt.Printf("textMes=%v, Type of textMes=%T \n", myu16, myu16)

}


// 结构体
type User struct{
	Name string `json:"name"`
	Id uint32
}
type Account struct{
	User
	password string
}
type Contact struct{
	 *User
	 Remark string
}

func Struct(){
	var user User = User{
		Name: "张三", 
	}
	user.Id = 1

	var user2 *User  = &User{
		Name: "李四",
	}
	user2.Id = 2

	 var a1 = Account{
		User: User{
			Name: "王五",
		},
		password: "123",
	 }

	 var c1 *Contact = &Contact{
		User: user2,
		Remark: "王麻子",
	 }

	 fmt.Println("a1 = ", a1)
	 fmt.Println("c1 = ", c1)



}

// 方法： 与特定类型绑定的函数
//类型的定义和方法需要在同一个包内
func (u User) printName(){
	fmt.Println("u.name = ", u.Name)
}

func (u *User) setId() {
	u.Id = 11
}
func Method() {
	u := User{
		Name: "A",
	}
	u.printName()
	u.setId()
	fmt.Println("u = ", u)
}

// 接口
// 接口本身不能绑定方法
// 接口是值类型，保存的是：值+原始类型 
type textMes struct{
	Type string
	Text string
}
func (tm *textMes) setText(){
	tm.Text = "TEXT"
}
type imgMes struct{
	Type string
	Img string
}
func (im *imgMes) setImg(){
	im.Img = "cat.jpg"
}

type Mes interface{
	setType()
}
func(tm *textMes)setType(){
	tm.Type = "Text Message"
}
func(im *imgMes)setType(){
	im.Type = "image Message"
}

func SendMes(m Mes){
	m.setType()

	switch mptr := m.(type){
	case *textMes:
		mptr.setText()
	case *imgMes:
		mptr.setImg()
	}

	fmt.Println("m = ", m)
}

func Interface(){
	tm :=textMes{}
	SendMes(&tm)
	im := imgMes{}
	SendMes(&im)

	var n1 int =1
	n1interface := interface{}(n1)

	n2, ok := n1interface.(int)
	if ok {
		fmt.Println("n2 = ", n2)
	}else {
		fmt.Println("类型断言失败")
	}

}

// 协程

var(
	c int
	lock sync.Mutex
)
func PrimeNum(n int){
	for i :=2; i < n; i++{
		 if n%i == 0 {
			return
		 }
	}
	fmt.Printf("%v \t", n)
	lock.Lock()
	c++
	lock.Unlock()

}

func Goroutine() {
	for i := 2; i < 100001; i++{
		go PrimeNum(i)
	}
	// var key string
	// fmt.Scanln(&key)
	fmt.Printf("\n找到%v个素数\n", c)
}

// channel 
func pushPrimeNum(n int, c chan int){
	for i :=2; i < n; i++{
		 if n%i == 0 {
			return
		 }
	}
	c <- n
	// close(c)

}

func pushNum(c chan int){
	for i := 0; i < 100; i++{
		c <- i
	}
	close(c)
}
func Channel() {
	var c1 chan int = make(chan int)
	go pushNum(c1)

	for v := range c1{
		fmt.Printf("%v \t", v)
	}

	var c2 chan int = make(chan int, 100)

	for i := 2; i < 100001; i++{
		go pushPrimeNum(i, c2)
	}
	Print:
	for{
		select{
		case v := <- c2:
			fmt.Printf("%v \t", v)
		default:
			fmt.Println("所有素数都找到了")
			break Print
		}

	}
}
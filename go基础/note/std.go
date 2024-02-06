package note

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"hello_go/util"
	"io"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// 随机数
func RandNum() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10) + 1)
}

// 字符串类型转换

func StrConv() {
	i1 := 123
	s1 := "baidu"
	s2 := fmt.Sprintf("%d@%s", i1, s1)
	fmt.Println("s2 = ", s2)

	var (
		i2 int
		s3 string
	)

	n, err := fmt.Sscanf(s2, "%d@%s", &i2, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功解析了%d个数据", n)

	s4 := strconv.FormatInt(123, 4)
	fmt.Println("s4 = ", s4)
	u1, err := strconv.ParseUint(s4, 4, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("u1 = ", u1)
}

// strings包常见函数
func PackageStrings() {
	fmt.Println(strings.Contains("hello", "o"))
	fmt.Println(strings.Replace("hello", "o", "ok", 1))
	fmt.Println(strings.Fields("hello hello \n hello"))
}

// 中文字符常见操作
func PackageUtf8() {
	str := "hello 世界"
	fmt.Println(utf8.ValidString(str))

}

// 时间常见操作
func PackageTime() {
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		fmt.Println()
		time.Sleep(time.Microsecond)
	}
	fmt.Println()

	d1, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println("d1 = ", d1)

	var intChan chan int = make(chan int)
	select {
	case <-intChan:
		fmt.Println("收到验证码")
	case <-time.After(time.Second):
		fmt.Println("验证码已过期")
	}

	fmt.Println()

	l1, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println(l1.String())

	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2001年1月1日， 12点45分"))

	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
TickerFor:
	for {
		select {
		case <-intChan:
			fmt.Println()
			break TickerFor
		case <-time.NewTicker(time.Microsecond).C:
			fmt.Print(".")
		}

	}

}

// 文件常见操作
func FileOperation() {
	// util.MkdirWithFilePath("A/B/V")
	// 文件夹操作
	dirEntrys, err := os.ReadDir("/Users/ckk/go")
	if err != nil {
		panic(err)
	}
	for _, v := range dirEntrys {
		fmt.Println(v.Name())
	}

	// 打开文件
	file, err := os.OpenFile("file1", os.O_RDWR|os.O_CREATE, 0665)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入文件

	data, err := os.ReadFile("file1")
	if err != nil {
		panic(err)
	}
	fmt.Println("data in file1:", string(data))
	err = os.WriteFile("file2", data, 0775)
	if err != nil {
		panic(err)
	}
}

// 文件读写
func FileReadAndWrite() {
	file, err := os.OpenFile("file5", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 1; i < 3; i++ {
		fileName := fmt.Sprintf("file%v", i)
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		data = append(data, '\n')
		writer.Write(data)
	}
	writer.Flush()
}

func Error() {
	defer func() {
		err := recover()
		fmt.Println("捕捉到了错误：", err)
	}()
	err1 := errors.New("Error One")
	fmt.Println("err1 = ", err1)
	err2 := fmt.Errorf("Error Two")
	fmt.Println("err2 = ", err2)
	panic(err1)
}

// 日志
func Log() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	err := errors.New("One error")
	util.INFO.Println(err)
	// util.WARN.Panicln(err)
	util.ERR.Fatalln(err)
}

// 单元测试
func IsNotNagative(n int) bool {
	return n > -1
}

// 命令行参数
func CmdArgs() {
	fmt.Printf("接收到了%v个参数\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("第%v个参数是%v \n", i, v)
	}
	fmt.Println()

	vPtr := flag.Bool("v", false, "GoNote版本")
	var userName string
	flag.StringVar(&userName, "u", "", "用户名")

	flag.Func("f", "", func(s string) error {
		fmt.Println("s = ", s)
		return nil
	})
	flag.Parse()
	if *vPtr {
		fmt.Println(*vPtr)
		fmt.Println("GoNote版本是 v0.0.0")
	}
	fmt.Println("输入用户名：", userName)

	for i, v := range flag.Args() {
		fmt.Printf("第%v个参数是%v \n", i, v)
	}

}

// runtime包
func PackageRuntime() {
	if runtime.NumCPU() > 7 {
		runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	}
	// runtime.Goexit()
}

// 同步
func PackageSync() {
	var c int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	primeNum := func(n int) {
		defer wg.Done()
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return
			}
		}
		mutex.Lock()
		c++
		mutex.Unlock()

	}
	for i := 2; i < 100001; i++ {
		wg.Add(1)
		go primeNum(i)
	}

	wg.Wait()
	fmt.Printf("找到%v个素数\n", c)

	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go func(n int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("协程%v被唤醒\n", n)
			cond.L.Unlock()
		}(i)
	}

	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		if i == 4 {
			fmt.Println()
			cond.Signal()

		}
		if i == 9 {
			fmt.Println()
			cond.Broadcast()
		}
	}

	fmt.Println()

	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("只执行一次")
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

// JSON常见操作
func PackageJson() {
	type user struct {
		Name string `json:"name"`
		Age  int
	}

	u1 := user{
		Name: "aa",
		Age:  1,
	}

	data, _ := json.Marshal(u1)

	fmt.Println(string(data))

	buf := new(bytes.Buffer)
	json.Indent(buf, data, "", "\t")
	fmt.Println(buf.String())

	var u2 user
	json.Unmarshal(data, &u2)
	fmt.Println("u2 = ", u2)
}

//	Tcp编程入门
func TcpCli(){
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil{
		fmt.Println("拨号失败")
		return
	}
	defer conn.Close()
	for{
		mes := struct {
			UserName string
			Mes string
		}{
			UserName: "kk",
		}
		fmt.Println("请输入要发送的内容：")
		fmt.Scanf("%s\n", &mes.Mes)
		if mes.Mes == "" {
			fmt.Println("输入为空")
			continue
		}
		if mes.Mes == "exit"{
			return
		}

		// data, _ := json.Marshal(&mes)
		// n, err := conn.Write(data)
		// if err != nil{
		// 	fmt.Println("发送失败")
		// 	return
		// }
		// fmt.Printf("发送成功字节%v\n", n)

		json.NewEncoder(conn).Encode(&mes)
		if err != nil {
			fmt.Println("发送失败")
			return
		}
	}
}

func TcpServer() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败：:", err)
		return
	}
	defer listener.Close()
	for{
		fmt.Println("主进程等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接听失败：", err)
			continue
		}
		go func(conn net.Conn){
			fmt.Println("一个客户端协程已开启")
			defer conn.Close()

			for {
				// buf := make([]byte, 4096)
				// n, err := conn.Read(buf)
				// if err == io.EOF {
				// 	fmt.Println("客户端退出了")
				// 	return
				// }
				// if err != nil {
				// 	fmt.Println("读取消息失败")
				// 	return
				// }
				// json.Unmarshal(buf[:n], &mes)
				mes := struct {
					UserName string
					Mes string
				}{}
				err = json.NewDecoder(conn).Decode(&mes)
				if err == io.EOF {
					fmt.Println("客户端退出了")
					return
				}
				if err != nil {
					fmt.Println("读取消息失败")
					return
				}
				fmt.Printf("%v说: %v\n", mes.UserName, mes.Mes)

			}
		}(conn)
	}

}

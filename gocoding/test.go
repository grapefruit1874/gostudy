package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// test1
type People struct {
	name string `json:"name"` //小写开头私有，不应该加json标签，也无法打印
}

func peopletest1() {
	js := `{
"name":"11"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("people:", p)
}

// test2
type t struct {
	n int
}

func gett() t {
	return t{} //直接返回等t{}无法寻址，无法直接赋值
}
func test2() {
	t := gett()
	p := &t.n
	*p = 1
	fmt.Println(t.n)
}
func test3() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x)
}

type stu struct {
	Name string
}

func test4() {
	fmt.Println(&stu{Name: "1"} == &stu{Name: "1"}) //指针类型比较的是指针地址
	fmt.Println(stu{Name: "1"} == stu{Name: "1"})   //非指针类型比较的是每个属性的值
}

func test5() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]                   //{"b","c"}
	str2[1] = "new"                    //{"b","new"}
	fmt.Println(str1)                  //{"a","b","new"}
	str2 = append(str2, "z", "x", "y") //导致底层数组扩容,会指向新的数组指针，追加的数据不影响str1
	fmt.Println(str1)                  //{"a","b","new"}
}

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

func test() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//:=不能用于结构体字段赋值
//变量隐藏
//自增，自减不在是运算符，只能作为独立语句，而不是表达式
//go语言不支持++i,--i
//nil用于表示interface,函数，map,slice,chan的零值，不能分配给string类型 	var str string=nil
//常量组中如果不知道类型和初始值，则与上一行非空常量右值相等
//用字面量初始化数组，slice,map 最好是在每个元素后面加上逗号，即使声明在一行或者多行都不会出错
//cap()适用于数组，数组指针，slice,chan不适用于map
//不能对nil的map直接赋值，需要使用make初始化，对nil的slice可以使用append()
//类型断言：i.(Type) i是接口，type是类型或接口???不太清楚怎么使用
//import _引用 但不使用
//有方向的chan不能被关闭
//使用=是定义类型别名 struct
//发生类型强转都会发生内存copy
//unsafe.pointer(&a)可以得到变量a的地址
//waitgroup在调用wait之前不能再调用add
//多核cpu,因为cpu缓存会导致多个核心中变量值不同步
//开锁后赋值变量，会将锁的状态也赋值，已加锁，再次加锁就会死锁
//数组只能与相同维度长度以及类型比较，切片之前不能直接比较

// 调度器优先调用哪个g,当创建一个g时，会优先放入到下一个调度的runnext作为下一次优先调用的g
func test7() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i1:", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i2:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func test9() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go getpanic()
}

func getpanic() {
	panic("hhhhhh")

}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	go testpanic()
	time.Sleep(time.Second)
}
func testpanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	testpanic1()
}
func testpanic1() {
	panic("我错了11")
}

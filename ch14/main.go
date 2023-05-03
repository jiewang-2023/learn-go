package main

import "fmt"

// 内存分配 new make
// 如果要对一个变量赋值，这个变量必须有对应的分配好的内存，这样才可以对这块内存操作，完成赋值的目的。
// 对于指针变量，如果没有分配内存，取值操作一样会报 nil 异常，因为没有可以操作的内存。

//new函数 根据传入的类型申请一块内存，然后返回指向这块内存的指针，指针指向的数据就是该类型的零值
//new 函数只用于分配内存，并且把内存清零，也就是返回一个指向对应类型零值的指针。new 函数一般用于需要显式地返回指针的情况，不是太常用。

// make 函数
// make 函数只用于 slice、chan 和 map 这三种内置类型的创建和初始化，因为这三种类型的结构比较复杂，比如 slice 要提前初始化好内部元素的类型，slice 的长度和容量等
type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {
	p := new(person)
	p.name = name
	p.age = age
	return p

}

func main() {

	var s string

	fmt.Printf("%p\n", &s)

	sp1 := new(string)
	fmt.Println("sp1 ", *sp1) //打印空字符串,也就是string的零值。

	pp := NewPerson("洛洛", 20)
	fmt.Println(pp)

	m := make(map[string]int, 10)
	fmt.Println(m)
}

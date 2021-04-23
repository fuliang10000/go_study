package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	//v1 := []rune{'\u4ed8', '\u4eae'}
	//v1 := "100"
	//v2, _ := strconv.Atoi(v1)
	//var b [3][3]int
	//a := [...]int{1, 2, 3}
	//a := [5]int{1: 3, 3: 7}
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//}
	//for _, v := range a {
	//	fmt.Println("Element of arr is", v)
	//}
	//var multi [9][9]string
	//for j := 0; j < 9; j++ {
	//	for i := 0; i < 9; i++ {
	//		n1 := i + 1
	//		n2 := j + 1
	//		if n1 < n2 { // 摒除重复的记录
	//			continue
	//		}
	//		multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
	//	}
	//}
	//fmt.Println(multi)
	//// 打印九九乘法表
	//for _, v1 := range multi {
	//	for _, v2 := range v1 {
	//		fmt.Printf("%-8s", v2) // 位宽为8，左对齐
	//	}
	//	fmt.Println()
	//}
	//months := [...]int8{1, 2, 3, 4, 5, 6, 7, 8}
	//q2 := months[3:7]
	//fmt.Println(q2)
	//mySlice1 := make([]int, 5, 10)
	//fmt.Println(len(mySlice1))
	//fmt.Println(cap(mySlice1))
	//appendSlice := []int{1, 2, 3, 4, 5}
	//newSlice := append(mySlice1, appendSlice...)
	//fmt.Println(newSlice)
	//keys := make([]string, 0)
	//testMap := map[string]int{
	//	"three": 3,
	//	"one":   1,
	//	"two":   2,
	//}
	//for k, _ := range testMap {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	//fmt.Println(keys)
	//a := 100
	//ptr := &a
	//b := *ptr
	//a = 200
	//fmt.Println(ptr)
	//fmt.Println(b)
	//switch {
	//case a <= 50:
	//	fmt.Println(50)
	//case a <= 150:
	//	fmt.Println(150)
	//case a <= 200:
	//	//fallthrough
	//case a <= 250:
	//	fmt.Println(250)
	//}
	//sum := 0
	//num := 1
	//for  num <= 100{
	//	sum += num
	//	num++
	//}
	//fmt.Println(sum)
	//slice := []int{1, 2, 3, 4, 5}
	//myfunc(1, 2, true, "dfdfdf", slice)
	//fmt.Println(addNum(11, "dfdf"))
	//func(a, b int) {
	//	fmt.Println(a + b)
	//}(11, 21)
	//add := func (a, b int) int  {
	//	return a + b
	//}
	//fmt.Println(add(1, 2))
	//c := multiply(a, b)
	//fmt.Printf("%d x %d = %d\n", a, b, c)
	//a := 2
	//b := 5
	//decorator1 := execTime(multiply1)
	//c := decorator1(a, b)
	//fmt.Printf("%d x %d = %d\n", a, b, c)
	//
	//a = 1
	//b = 4
	//decorator2 := execTime(multiply2)
	//c = decorator2(a, b)
	//fmt.Printf("%d | %d = %d\n", a, b, c)

	//d := 50
	//decorator3 := execTime1(fibonacci2)
	//e := decorator3(d)
	//fmt.Printf("%d = %d\n", d, e)
	//decorator4 := execTime1(fibonacci3)
	//e = decorator4(d)
	//fmt.Printf("%d = %d\n", d, e)
	//var users
}
//func myfunc(numbers ...interface{}) {
//	for _, number := range numbers {
//		fmt.Println(number)
//	}
//}
// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int
type MultiPlyFunc1 func(int) int

// 乘法运算函数
func multiply1(a, b int) int {
	time.Sleep(1)
	return a * b
}
// 位运算函数
func multiply2(a, b int) int {
	time.Sleep(1)
	return a | b
}

func execTime(f MultiPlyFunc) MultiPlyFunc  {
	return func(a, b int) int {
		start := time.Now() // 起始时间
		c := f(a, b)
		end := time.Since(start)
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c
	}
}

func execTime1(f MultiPlyFunc1) MultiPlyFunc1  {
	return func(a int) int {
		start := time.Now() // 起始时间
		c := f(a)
		end := time.Since(start)
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c
	}
}

func fibonacci1(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return fibonacci1(n - 1) + fibonacci1(n - 2)
}
const MAX = 50
var fibs [MAX]int
func fibonacci2(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}
	num := fibonacci2(n-1) + fibonacci2(n-2)
	fibs[index] = num
	return num
}

func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}
	return fibonacciTail(n-1, second, first+second) // 48 1 2 48
}
func fibonacci3(n int) int {
	return fibonacciTail(n, 0, 1) // F(1) = 0, F(2) = 1
}

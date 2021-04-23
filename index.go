package main

import "fmt"

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
	slice := []int{1, 2, 3, 4, 5}
	myfunc(1, 2, true, "dfdfdf", slice)
	fmt.Println(addNum(11, "dfdf"))
	func(a, b int) {
		fmt.Println(a + b)
	}(11, 21)
}
func myfunc(numbers ...interface{}) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func addNum(a int, b string) (int, string) {
	return a, b
}

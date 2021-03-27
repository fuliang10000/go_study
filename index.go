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
	months := [...]int8{1, 2, 3, 4, 5, 6, 7, 8}
	q2 := months[3:7]
	fmt.Println(q2)
	mySlice1 := make([]int, 5)
	fmt.Println(mySlice1)
}

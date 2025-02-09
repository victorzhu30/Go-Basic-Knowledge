package main

import "fmt"

func modify1(array [5]string) {
	array[2] = "十"
	array[3] = "王"
}

func modify2(slice []string) {
	slice[2] = "地"
	slice[3] = "衡"
}

func modify3(m map[string]int) {
	m["chong"] = 100
}

func main() {

	//数组
	array := [5]string{"符", "玄", "太", "卜", "司"}
	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(cap(array))

	array2 := array
	array2[0] = "青"
	array2[1] = "雀"
	fmt.Println(array2)

	modify1(array)
	fmt.Println(array)

	//切片
	slice := []string{"符", "玄", "太", "卜", "司"}
	fmt.Println(slice, len(slice), cap(slice))

	slice2 := make([]int, 2)
	slice3 := make([]string, 2)
	fmt.Println(slice2, slice3)

	slice4 := make([]int, 5, 10)
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice5 := make([]int, 2, 2)
	slice6 := append(slice5, 1)
	fmt.Println(slice5, len(slice5), cap(slice5), slice6, len(slice6), cap(slice6))

	modify2(slice)
	fmt.Println(slice)

	//切片的切片
	subslice := slice[1:3] // 左闭右开
	fmt.Println(subslice)
	subslice[0] = "天"
	fmt.Println(subslice, slice)

	multiSlice := make([][]int, 0)
	multiSlice = append(multiSlice, []int{1, 2, 3})
	multiSlice = append(multiSlice, []int{4, 5, 6})

	fmt.Println(multiSlice)

	multiSlice[0][1] = 100
	fmt.Println(multiSlice)

	//映射
	scores := map[string]int{
		"ming":     10,
		"zhangsan": 13,
	}
	fmt.Println(scores, len(scores))
	fmt.Println(scores["ming"])

	score1, exist1 := scores["ming"]
	fmt.Println(score1, exist1)
	score2, exist2 := scores["chong"]
	fmt.Println(score2, exist2)

	modify3(scores)
	fmt.Println(scores)

	delete(scores, "ming")
	fmt.Println(scores)

	for key, value := range scores {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}

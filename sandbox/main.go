package main

import "fmt"

// ジェネリクスを使用して関数を一つにまとめる
func getArray[T any](items []T) {
	fmt.Println(items)
}

func main() {
	getArray([]string{"apple", "banana"}) // ["apple", "banana"]
	getArray([]int{1, 2, 3})              // [1, 2, 3]
}

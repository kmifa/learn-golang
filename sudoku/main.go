package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Boardはsuudokuのboardを示す
type Board [9][9]int

func pretty(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func duplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Board) bool {
	// 行チェック
	for i := 0; i < 9; i++ {
		// 出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
			// fmt.Println(c)
		}
		if duplicated(c) {
			return false
		}
	}

	// 列チェック
	for i := 0; i < 9; i++ {
		// 出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
			// fmt.Println(c)
		}
		if duplicated(c) {
			return false
		}
	}

	// 3x3チェック
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}

	return true
}

func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board) bool {
	// 実行内容を見たいときは以下のコメントを外す
	// time.Sleep(time.Second * 1)
	// fmt.Printf("%v\n", pretty(*b))
	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// ますが0だったら探索開始
			if b[i][j] == 0 {
				for c := 9; c >= 1; c-- {
					b[i][j] = c
					// もし数字がルールに適合するなら
					if verify(*b) {
						// さらに深く探索する
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {

	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Printf("%+v\n", pretty(b))
}

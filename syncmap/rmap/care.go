package rmap

import "fmt"

/*
注意 如果使用 struct{} 作为key，那么如果struct变化会影响map取值
*/

func care1() {
	type st struct {
		val int
	}
	m1 := make(map[st]int)
	s := st{val: 1}
	m1[s] = 1
	fmt.Println(m1[s])
	s.val = 2
	/*
		找不到 打印0
	*/
	fmt.Println(m1[s])
}

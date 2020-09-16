package practice

import "fmt"

func fun() bool {
	arr := [10]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5
	if len(arr) == 1 {
		return false
	}
	// remainder count
	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		m[arr[i]%k-1] = arr[i]
	}
	fmt.Println(m)
	return true
}

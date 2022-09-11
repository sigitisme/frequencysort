package main

import (
	"fmt"
	"sort"
)

func main() {

	inputs := []int32{3, 1, 2, 2, 4}

	items := itemsSort(inputs)

	for _, v := range items {
		fmt.Println(v)
	}
}

func itemsSort(items []int32) []int32 {
	m := make(map[int32]int)
	maxFreq := 0

	for _, v := range items {
		if val, ok := m[v]; !ok {
			m[v] = 0
		} else {
			m[v] = val + 1
			if maxFreq < m[v] {
				maxFreq = m[v]
			}
		}
	}

	result := make([][]int32, maxFreq+1)

	for k, v := range m {
		result[v] = append(result[v], k)
	}

	results := make([]int32, 0)

	for v := range result {
		sort.Slice(result[v], func(i, j int) bool {
			return result[v][i] < result[v][j]
		})

		results = append(results, result[v]...)
	}

	return results
}

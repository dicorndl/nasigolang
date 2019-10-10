package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	for i := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
		} else {
			m := merged[len(merged)-1][1]
			if m < intervals[i][1] {
				m = intervals[i][1]
			}
			merged[len(merged)-1][1] = m
		}
	}

	return merged
}

func main() {
	a := [][]int{
		{1, 3},
		{8, 10},
		{15, 18},
		{2, 6},
	}
	b := merge(a)

	fmt.Println(b)
}

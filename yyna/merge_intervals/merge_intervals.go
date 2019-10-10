import (
  "fmt"
  "sort"
)

type matrix [][]int

func (m matrix) Len() int {
  return len(m)
}

func (m matrix) Swap(i, j int) {
  m[i], m[j] = m[j], m[i]
}

func (m matrix) Less(i, j int) bool {
  return m[i][0] < m[j][0]
}

func merge(intervals [][]int) [][]int {
  if intervals == nil || len(intervals) < 1 {
    return intervals
  }

  result := [][]int{}
  sort.Sort(matrix(intervals))

  var curStart, curEnd, nextStart, nextEnd int
  curStart, curEnd = intervals[0][0], intervals[0][1]

  for i := 1; i < len(intervals); i++ {
    nextStart, nextEnd = intervals[i][0], intervals[i][1]

    if curEnd >= nextStart {
      curStart, curEnd = min(curStart, nextStart), max(curEnd, nextEnd)
    } else {
      result = append(result, []int{curStart, curEnd})
      curStart, curEnd = nextStart, nextEnd
    }
  }

  result = append(result, []int{curStart, curEnd})

  return result
}

func max(x, y int) int {
  if x < y {
    return y
  }
  return x
}

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

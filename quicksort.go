package main
import (
  "fmt"
  "time"
)

type QuickSort struct{}

func (qs QuickSort) partition(arr []int, l int, r int) int {
  // Index of smaller element and indicates the right position of pivot found so far
  i := l - 1

  for j := l; j < r; j++ {
    if arr[j] <= arr[r] {
      i++
      arr[i], arr[j] = arr[j], arr[i]
    }
  }
  arr[i+1], arr[r] = arr[r], arr[i+1]
  return i+1
}

func (qs QuickSort) sort(arr []int, l int, r int) {
  if l < r {
    p := qs.partition(arr, l, r)

    qs.sort(arr, l, p-1)
    qs.sort(arr, p+1, r)
  }
}

func main() {
  arr := []int{5,2,6,7,1,2,0,-1,10,99,1,50}
  s := QuickSort{}
  
  start := time.Now()
  defer func() {
    fmt.Println("exec time: ", time.Now().Sub(start))
  }()

  s.sort(arr, 0, len(arr)-1)
  fmt.Println(arr)
}

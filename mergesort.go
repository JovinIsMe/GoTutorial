package main
import (
  "fmt"
  "time"
)

type MergeSort struct{}

func (ms MergeSort) merge(arr []int, l int, m int, r int) {
  // size of left and right array
  n1 := m - l + 1
  n2 := r - m

  // copy left and right array
  arr1 := make([]int, n1)
  copy(arr1,arr[l:m+1])
  arr2 := make([]int, n2)
  copy(arr2,arr[m+1:r+1])

  // change the original array with sorted values
  i,j,k := 0,0,l
  for i<n1 && j<n2 {
    if arr1[i] < arr2[j] {
      arr[k] = arr1[i]
      i++
    } else {
      arr[k] = arr2[j]
      j++
    }
    k++
  }

  // copy the other values to original array
  for i < n1 {
    arr[k] = arr1[i]
    i++
    k++
  }
}

func (ms MergeSort) sort(arr []int, l int, r int) {
  if l < r {
    m := l + (r-l)/2
    ms.sort(arr, l, m)
    ms.sort(arr, m+1, r)
    ms.merge(arr, l, m, r)
  }
}

func main() {
  arr := []int{5,2,6,7,1,2,0,-1,10,99,1,50}
  s := MergeSort{}
  
  start := time.Now()
  defer func() {
    fmt.Println("exec time: ", time.Now().Sub(start))
  }()

  s.sort(arr, 0, len(arr)-1)
  fmt.Println(arr)
}

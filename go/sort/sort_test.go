package sort

import (
    "fmt"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    nums := []int{10, 8, 12, 5, 11, 3}
    bubbleSort(nums)
    fmt.Println(nums)
}

func TestQuickSort(t *testing.T) {
    nums := []int{15, 10, 34, 22, 1, 0, 3, 8}
    quickSort(0, len(nums)-1, nums)
    fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
    nums := []int{15, 10, 34, 22, 1, 0, 3, 9}
    mergeSort(0, len(nums)-1, nums)
    fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
    nums := []int{12, 10, 23, 99, 23, 18, 34, 1}
    heapSort(nums)
    fmt.Println(nums)
}

package sort

// 冒泡排序 
func bubbleSort(nums []int) {
    length := len(nums)
    if length == 0 { return }
    for cmpCnt := 0; cmpCnt < length-1; cmpCnt ++ {
        for item := 1; item < length-cmpCnt; item++ {
            if nums[item] < nums[item-1] {
                nums[item], nums[item-1] = nums[item-1], nums[item]
            }
        }
    }
}


// 快排
func quickSort(start, end int, nums []int) {
    if start < end {
        aim := nums[start]
        s, e := start, end
        for s < e {
            for s < e && nums[s] < aim {s++}
            for s < e && nums[e] > aim {e--}
            if s < e {
                nums[s], nums[e] = nums[e], nums[s]
            }
        }
        quickSort(start, s-1, nums)
        quickSort(s+1, end, nums)
    }
}


// 归并排序
func merge(start, end int, nums []int) {
    mid := (start + end)/2
    i, j := start, mid+1
    newArray := make([]int, 0)
    for i <= mid && j <= end {
        if nums[i] < nums[j] {
            newArray = append(newArray, nums[i])
            i++
        } else {
            newArray = append(newArray, nums[j])
            j++
        }
    }
    for ; i <= mid; i++ {
        newArray = append(newArray, nums[i])
    }
    for ; j <= end; j++ {
        newArray = append(newArray, nums[j])
    }
    for index := 0; index < len(newArray); index++ {
        nums[start+index] = newArray[index]
    }
}

func mergeSort(start, end int, nums []int) {
    if start < end {
        mid := (start + end)/2
        mergeSort(start, mid, nums)
        mergeSort(mid+1, end, nums)
        merge(start, end, nums)
    }
}


// 堆排序
func maintain(nums []int, index, heapSize int) {
    left, right := index*2+1, (index+1)*2
    aimIndex := index
    if left <= heapSize && nums[aimIndex] < nums[left] {
        aimIndex = left
    }
    if right <= heapSize && nums[aimIndex] < nums[right] {
        aimIndex = right
    }
    if aimIndex != index {
        nums[index], nums[aimIndex] = nums[aimIndex], nums[index]
        maintain(nums, aimIndex, heapSize)
    }
}

func buildHeap(nums []int) {
    length := len(nums)
    for index := length/2; index >= 0; index-- {
        maintain(nums, index, length-1)
    }
}

func heapSort(nums []int) {
    buildHeap(nums)
    length := len(nums)-1
    for length > 0 {
        nums[0], nums[length] = nums[length], nums[0]
        length--
        maintain(nums, 0, length)
    }
}

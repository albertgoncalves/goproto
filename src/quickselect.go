package main

import (
    "fmt"
    "math/rand"
    "time"
)

type t = uint8

func partition(array []t, left, right, pivotIndex int) int {
    pivotValue := array[pivotIndex]
    array[pivotIndex], array[right] = array[right], array[pivotIndex]
    storeIndex := left
    for i := left; i < right; i++ {
        if array[i] < pivotValue {
            array[storeIndex], array[i] = array[i], array[storeIndex]
            storeIndex++
        }
    }
    array[storeIndex], array[right] = array[right], array[storeIndex]
    return storeIndex
}

func quickSelect(r *rand.Rand, array []t, left, right, k int) t {
    for {
        if left == right {
            return array[left]
        }
        pivotIndex := partition(array, left, right, left+r.Intn(right-left+1))
        if k == pivotIndex {
            return array[k]
        } else if k < pivotIndex {
            right = pivotIndex - 1
        } else {
            left = pivotIndex + 1
        }
    }
}

func quickSort(r *rand.Rand, array []t, left, right int) {
    if left < right {
        pivotIndex := partition(
            array,
            left,
            right,
            left+r.Intn(right-left+1),
        )
        quickSort(r, array, left, pivotIndex)
        quickSort(r, array, pivotIndex+1, right)
    }
}

func main() {
    array := []t{3, 4, 2, 5, 7, 1, 8, 2, 9, 10}
    n := len(array)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    fmt.Println("\noriginal array   :", array)
    {
        for i := 0; i < n; i++ {
            a := make([]t, len(array))
            copy(a, array)
            v := quickSelect(r, a, 0, n-1, i)
            fmt.Println("post-quickselect :", a, "| k :", i, "| value :", v)
        }
    }
    {
        a := make([]t, len(array))
        copy(a, array)
        quickSort(r, a, 0, n-1)
        fmt.Println("post-quicksort   :", a)
    }
}

package main

import (
    "cmp"
    "fmt"
    "math/rand"
    "time"
)

func partition[T cmp.Ordered](array []T, l, r, i int) int {
    x := array[i]

    array[i], array[r] = array[r], array[i]

    j := l
    for i := l; i < r; i++ {
        if array[i] < x {
            array[j], array[i] = array[i], array[j]
            j++
        }
    }

    array[j], array[r] = array[r], array[j]

    return j
}

func quickSelect[T cmp.Ordered](rng *rand.Rand, array []T, l, r, k int) T {
    for {
        if l == r {
            return array[l]
        }

        i := partition(array, l, r, l+rng.Intn(r-l+1))

        if k == i {
            return array[k]
        } else if k < i {
            r = i - 1
        } else {
            l = i + 1
        }
    }
}

func quickSort[T cmp.Ordered](rng *rand.Rand, array []T, l, r int) {
    if r <= l {
        return
    }

    i := partition(array, l, r, l+rng.Intn(r-l+1))
    quickSort(rng, array, l, i)
    quickSort(rng, array, i+1, r)
}

func main() {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < 10; i++ {
        array := []int{3, 4, 2, 5, 7, 1, 8, 2, 9, 10}
        v := quickSelect(rng, array, 0, len(array)-1, i)
        fmt.Println("post-quickselect :", array, "| k :", i, "| value :", v)
    }

    array := []int{3, 4, 2, 5, 7, 1, 8, 2, 9, 10}
    quickSort(rng, array, 0, len(array)-1)
    fmt.Println("post-quicksort   :", array)
}

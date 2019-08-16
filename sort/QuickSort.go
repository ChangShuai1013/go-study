package main

import "strconv"

func QuickSort(arr *[]int, p int, r int) {
	if p < r {
		q := Partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}

func Partition(arr *[]int, p int, r int) int {
	x := (*arr)[r]
	i := p - 1
	for j := p; j < r; j++ {
		if (*arr)[j] <= x {
			i++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = tmp
		}
	}
	tmp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[r]
	(*arr)[r] = tmp
	return i + 1
}

func main() {
	arr := make([]int, 0)
	arr = append(arr, 2)
	arr = append(arr, 8)
	arr = append(arr, 7)
	arr = append(arr, 1)
	arr = append(arr, 3)
	arr = append(arr, 5)
	arr = append(arr, 6)
	arr = append(arr, 4)
	QuickSort(&arr, 0, len(arr)-1)
	for _, num := range arr {
		print(strconv.Itoa(num) + "\t")
	}
	println()
}

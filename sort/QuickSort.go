package main

func QuickSort(arr *[]int, p int, r int) {
	if p < r {
		q := Partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}

func Partition(arr *[]int, p int, r int) int {
	x := *arr[r]

}

func main() {

}

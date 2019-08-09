package main

func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)
	indexs := make([]int, 0)
	for i, num := range nums {
		complement := target - num
		if _, ok := numIndex[complement]; ok {
			indexs = append(indexs, i)
			indexs = append(indexs, numIndex[complement])
			return indexs
		}
		numIndex[num] = i
	}
	return nil
}

func main() {
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	result := twoSum(nums, 7)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}

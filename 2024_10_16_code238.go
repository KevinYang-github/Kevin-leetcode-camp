package main

import "fmt"

func productExceptSelf(nums []int) []int {
	N := len(nums)
	answer := make([]int, N)
    answer[0] = 1
	tem := 1
	//先计算所有元素的左侧乘积，最左侧元素的左侧乘积为1，所以索引从1开始即可
	for i:=1; i<N; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	//下一步从后往前遍历，并使用tem保存上一个元素的右侧乘积，然后与当前的左侧乘积相乘写入answer中
	//最右侧的元素的右侧乘积为1，即其左侧乘积为该元素的答案。
	for i:=N-2; i>=0; i-- {
		tem *= nums[i+1]
		answer[i] *= tem
	}

	return answer
}

func main() {
	nums := []int{1,2,3,4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
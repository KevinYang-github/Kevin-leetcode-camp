package main

import "math"

func firstMissingPositive(nums []int) int {
    N := len(nums)

    for i:=0; i < N; i++ {
        if nums[i] <= 0 {
            nums[i] = N + 1
        }
    }

    for i:=0; i < N; i++ {
        abs_n := int(math.Abs(float64(nums[i])))
        if abs_n <= N{
            nums[abs_n - 1] = - int(math.Abs(float64(nums[abs_n - 1])))
        }
    }
    for i:=0; i < N; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    return N + 1
}

func main() {
	nums := []int{1,2,3,4}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
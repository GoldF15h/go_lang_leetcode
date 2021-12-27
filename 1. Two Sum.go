package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    var op []int
    for i := 0 ; i < len(nums) - 1 ; i++ {
        // fmt.Println(nums[i])
        for j := i+1 ; j < len(nums) ; j++{
            if nums[i] + nums[j] == target {
                return []int{i,j}
            }
        }
    }
    return op
}

func main() {
    // [2,7,11,15]
    // 9
    nums := []int{2,7,11,15}
    var tar int = 9
    fmt.Println(twoSum(nums,tar))
    
}
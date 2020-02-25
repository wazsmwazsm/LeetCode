// see: https://leetcode.com/problems/two-sum/description/

package main

import "fmt"

func twoSum(nums []int, target int) []int {

    if len(nums) < 2 {
      return nil
    }

    m := make(map[int]int, len(nums))

    for i, v := range nums {
        // 找到了满足条件的另一个值时返回
        if j, ok := m[v]; ok {
            return []int{j, i}
        }
        // 构建映射 (满足条件的一个值作为 key , 另一个作为建)
        m[target - v] = i
    }

    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 22

    fmt.Println(twoSum(nums, target))
}

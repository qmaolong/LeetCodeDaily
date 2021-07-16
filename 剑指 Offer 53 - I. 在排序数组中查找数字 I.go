package main

import "errors"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	position, err := findTarget(nums, target, left, right)
	if err != nil {
		return 0
	}
	total := 1
	i := 1
	for {
		find := false
		if position-i >= 0 && nums[position-i] == target {
			total++
			find = true
		}
		if position+i < len(nums) && nums[position+i] == target {
			total++
			find = true
		}
		if !find {
			break
		}
		i++
	}
	return total
}

func findTarget(nums []int, target int, left int, right int) (int, error) {
	if right-left <= 1 {
		if nums[left] == target {
			return left, nil
		} else if nums[right] == target {
			return right, nil
		}
		return 0, errors.New("")
	}
	middle := (left + right) / 2
	if nums[middle] == target {
		return middle, nil
	} else if nums[middle] > target {
		right = middle
	} else {
		left = middle
	}
	return findTarget(nums, target, left, right)

}

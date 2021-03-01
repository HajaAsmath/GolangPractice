package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(numbers, 11, 0, len(numbers)-1)
	if result {
		fmt.Println("Key found")
	} else {
		fmt.Println("Key not found")
	}
	if binarySearchRec(0, len(numbers)-1, 2, numbers) {
		fmt.Println("Key Found")
	} else {
		fmt.Println("Key not Found")
	}
}

func binarySearch(number []int, key int, start int, end int) bool {
	for start <= end {
		mid := int(start + (end-1)/2)
		fmt.Println(mid)
		if number[mid] == key {
			return true
		}
		if number[mid] < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func binarySearchRec(min int, max int, key int, numbers []int) bool {
	if min <= max {
		mid := int(min + (max-1)/2)
		if numbers[mid] == key {
			return true
		}
		if numbers[mid] < key {
			return binarySearchRec(mid+1, max, key, numbers)
		} else {
			return binarySearchRec(min, mid-1, key, numbers)
		}
	}
	return false
}

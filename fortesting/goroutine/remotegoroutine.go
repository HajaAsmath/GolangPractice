package main

import "fmt"

func sort(a []int, n int) {
	i := 0
	j := i + 1
	temp := 0
	for j < n {
		if a[i]%2 == 0 && a[j]%2 == 0 {
			i++
			j++
		} else if a[i]%2 != 0 && a[j]%2 != 0 {
			j++
		} else if a[i]%2 == 0 && a[j]%2 != 0 {
			i++
			j++
		} else if a[i]%2 != 0 && a[j]%2 == 0 {
			temp = a[i]
			a[i] = a[j]
			a[j] = temp
			i++
			j++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i], ",")
	}

}

func main() {
	a := []int{2, 3, 9, 8, 7, 5, 6, 4}
	sort(a, len(a))
}

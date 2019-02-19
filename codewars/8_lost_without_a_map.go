package main

// my
func Maps(x []int) (result []int) {
	for _, value := range x {
		result = append(result, value*2)
	}
	return
}

// best
func Maps(x []int) []int {
	var res = []int{}
	for _, num := range x {
		res = append(res, num*2)
	}
	return res
}

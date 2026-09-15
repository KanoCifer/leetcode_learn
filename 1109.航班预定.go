package leetcodelearn


func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		diff[booking[0]-1] += booking[2]
	    diff[booking[1]] -= booking[2]
	}

	result := make([]int, n)
	current := 0
	for i := 0; i < n; i++ {
		current += diff[i]
		result[i] = current
	}
	return result
}
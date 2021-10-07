package main

func absInt(x float32) float32 {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y float32) float32 {
	if x < y {
		return y - x
	}
	return x - y
}

// Convert integer to bool for check in 'if' condition
func itob(x int) bool {
	if x == 0 {
		return  true
	}
	return false
}


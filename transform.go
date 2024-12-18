package gostl

func Fill[T any](arr []T, val T) {
	if len(arr) == 0 {
		return
	}
	arr[0] = val
	for idx := 1; idx < len(arr); idx <<= 1 {
		copy(arr[idx:], arr[:idx])
	}
}

func FillZero[T any](arr []T) {
	var zero T
	Fill(arr, zero)
}

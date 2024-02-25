package sort

func BubbleSort(arr []int) {
	// Takes an array and sorts the array in place
	// Running time O(n^2)
	for arrLen := len(arr); arrLen > 0; arrLen-- {
		for i := 0; i < arrLen-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1] // swap
			}
		}
	}
}

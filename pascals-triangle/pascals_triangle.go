package pascal

func Triangle(n int) (t [][]int) {
	t = make([][]int, n)
	t[0] = []int{1}
	for i := 1; i < n; i++ {
		arr := make([]int, i+1)
		arr[0], arr[i] = 1, 1
		for j := 1; j < i; j++ {
			arr[j] = t[i-1][j-1] + t[i-1][j]
		}
		t[i] = arr
	}

	return t
}

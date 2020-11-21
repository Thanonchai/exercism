package pascal

func Triangle(n int) (t [][]int) {
	t = make([][]int, n)
	t[0] = []int{1}
	for i := 1; i < n; i++ {
		arr := make([]int, i+1)
		for j := 0; j <= i; j++ {
			left, right := 0, 0
			if j-1 >= 0 {
				left = t[i-1][j-1]
			}
			if j < i {
				right = t[i-1][j]
			}
			arr[j] = left + right
		}
		t[i] = arr
	}

	return t
}

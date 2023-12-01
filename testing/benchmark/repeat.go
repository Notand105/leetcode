package benchmark

func Repeat(str string, n int) string {
	aux := ""

	for i := 0; i < n; i++ {
		aux += str
	}

	return aux
}

package integers

func Add(a, b int) int {
	return a + b
}

// sumar los numero ambos postivos y ambos negativos, retornar un arreglo de ellos
func DoubleAdd(a, b int) []int {
	sum := a + b
	nsum := sum * -1
	return []int{sum, nsum}
}

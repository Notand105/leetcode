package slices

func SumAll(nums []int) int {
	var aux int = 0
	for i := 0; i < len(nums); i++ {
		aux += nums[i]
	}
	// we can do the same with range
	//for _, number := range nums{
	//	sum += number
	//}
	return aux
}

// recibir n cantidad de slices de m elementos y retornar un slices de n elementos
// conteniendo la suma de cada slice recibido

func SumMultiple(nums ...[]int) []int {
	length := len(nums)
	//make permite crear slices con una capacidad predeterminada
	sums := make([]int, length)

	for i, numbers := range nums {
		sums[i] = SumAll(numbers)
	}
	return sums
}

// lo mismo que la anterior que solo sumando las colas, ignorando las cabezas
func SumMultipleTails(nums ...[]int) []int {
	length := len(nums)
	//make permite crear slices con una capacidad predeterminada
	sums := make([]int, length)
	//Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the sides
	//of the : it captures everything to that side of it. In our case,
	//we are saying "take from 1 to the end" with numbers[1:]
	for i, numbers := range nums {
		tail := numbers[1:]
		sums[i] = SumAll(tail)
	}
	return sums
}

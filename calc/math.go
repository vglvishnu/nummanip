package calc

//Returns sum of two integers
func Add(i ...int) int {
	sum := 0
	for _, num := range i {
		sum = sum + num
	}
	return sum
}

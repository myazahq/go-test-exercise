package slices

func SumOfEvenNumbers(nums []int) int {
	var sum int
	if len(nums) == 0 {
		return sum
	}
	for _, num := range nums{
		if num % 2 == 0{
			sum += num
		}
	}
	return sum 
}
// package exampletest asks if you are ready to rock
package exampletest

func Sum(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

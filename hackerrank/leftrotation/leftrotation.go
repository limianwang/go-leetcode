/*
A left rotation operation on an array of size  shifts each of the array's elements  unit to the left. For example, if  left rotations are performed on array , then the array would become .

Given an array of  integers and a number, , perform  left rotations on the array. Then print the updated array as a single line of space-separated integers.
*/

package leftrotation

func rotateArray(arr []int, k int) []int {
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		index := (i + k) % len(arr)
		result = append(result, arr[index])
	}

	return result
}

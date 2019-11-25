package main

func minimumSwaps(arr []int32) int32 {
	arrLength := int32(len(arr))
	min := findMin(arr)

	var i int32
	var swapIndex int32
	var swaps int32

	for true {
		if arr[i] != min+i {
			swapIndex = arr[i]
			arr[i] = arr[swapIndex-min]
			arr[swapIndex-min] = swapIndex
			swaps++
		} else {
			i++
		}
		if i == arrLength-1 && arr[arrLength-1] == min+arrLength-1 {
			break
		}
	}

	return swaps
}

func findMin(arr []int32) int32 {
	var min = arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
	}

	return min
}

func main() {}

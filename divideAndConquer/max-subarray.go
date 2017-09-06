package divideAndConquer

// MaxSubArrayRec is the reference implementation for the
// maximum sub-array Problem im CLRS.
// It finds the subarray with the maximum sum in O(n lg n) time.
func MaxSubArrayRec(array []int) ([]int, int) {
	low, high, sum := maxSubArray(array, 0, len(array))
	return array[low : high+1], sum
}

func maxSubArray(array []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, array[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := maxSubArray(array, low, mid)
	rightLow, rightHigh, rightSum := maxSubArray(array, mid+1, high)
	crossLow, crossHigh, crossSum := maxCrossingSubArray(array, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}
func maxCrossingSubArray(array []int, low, mid, high int) (int, int, int) {
	leftSum := -1 << 31
	sum := 0
	maxLeft := 0
	for i := mid; i >= low; i-- {
		sum += array[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := -1 << 31
	sum = 0
	maxRight := 0
	for j := mid + 1; j <= high; j++ {
		sum += array[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

// MaxSubArrayLin is an implementation for the max sum-array
// problem in linear time.
// Solution to exercise 4.1-5
func MaxSubArrayLin(array []int) ([]int, int) {
	low, high := 0, 0
	negativeSum := 0
	if array[0] < 0 {
		negativeSum = array[0]
	}
	sum := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > 0 {
			if high == i-1 {
				high = i
				sum += array[i]
			} else if array[i]+negativeSum > 0 {
				high = i
				sum += array[i] + negativeSum
			} else {
				sum = array[i]
				low, high = i, i
				negativeSum = 0
			}
		} else if array[i] > sum {
			sum = array[i]
			low, high = i, i
			negativeSum = 0
		} else if array[i] <= 0 {
			negativeSum += array[i]
		}
	}
	return array[low : high+1], sum
}

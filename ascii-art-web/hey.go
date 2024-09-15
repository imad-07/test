package pushswap

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

// SwapA swaps the first two elements of stack A.
// If the stack has fewer than 2 elements, it does nothing.
func SwapA(a, b *[]int) {
	if len(*a) < 2 {
		return
	}
	(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
}

// SwapB swaps the first two elements of stack B.
// It calls SwapA with the arguments reversed.
func SwapB(a, b *[]int) {
	SwapA(b, a)
}

// SwapBoth swaps the top elements of both stack A and stack B.
func SwapBoth(a, b *[]int) {
	SwapA(a, b)
	SwapB(a, b)
}

// PushA pushes the top element from stack B to stack A.
// If stack B is empty, it does nothing.
func PushA(a, b *[]int) {
	if len(*b) == 0 {
		return
	}
	*a = append([]int{(*b)[0]}, *a...)
	*b = (*b)[1:]
}

// PushB pushes the top element from stack A to stack B.
// It calls PushA with the arguments reversed.
func PushB(a, b *[]int) {
	PushA(b, a)
}

// RotateA rotates stack A upwards by one position.
// The first element becomes the last one.
func RotateA(a, b *[]int) {
	if len(*a) == 0 {
		return
	}
	first := (*a)[0]
	*a = append((*a)[1:], first)
}

// RotateB rotates stack B upwards by one position.
// It calls RotateA with the arguments reversed.
func RotateB(a, b *[]int) {
	RotateA(b, a)
}

// RotateBoth rotates both stack A and stack B upwards by one position.
func RotateBoth(a, b *[]int) {
	RotateA(a, b)
	RotateB(a, b)
}

// ReverseRotateA rotates stack A downwards by one position.
// The last element becomes the first one.
func ReverseRotateA(a, b *[]int) {
	if len(*a) == 0 {
		return
	}
	last := (*a)[len(*a)-1]
	*a = append([]int{last}, (*a)[:len(*a)-1]...)
}

// ReverseRotateB rotates stack B downwards by one position.
// It calls ReverseRotateA with the arguments reversed.
func ReverseRotateB(a, b *[]int) {
	ReverseRotateA(b, a)
}

// ReverseRotateBoth rotates both stack A and stack B downwards by one position.
func ReverseRotateBoth(a, b *[]int) {
	ReverseRotateA(a, b)
	ReverseRotateB(a, b)
}

// CreateStackA creates stack A from a space-separated string of integers.
// It returns the created stack and an error string if any issues occur.
func CreateStackA(args string) ([]int, string) {
	var stack []int
	seen := make(map[int]bool)
	x := strings.Split(args, " ")
	for _, arg := range x {
		if arg == "" {
			continue
		}
		number, err := strconv.Atoi(string(arg))
		if err != nil {
			return nil, "Error"
		}
		if seen[number] {
			return nil, "Error"
		}
		seen[number] = true
		stack = append(stack, number)
	}
	return stack, ""
}

// IsSorted checks if the given stack is sorted in ascending order.
func IsSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] > stack[i] {
			return false
		}
	}
	return true
}

// Min returns the index of the minimum element in the stack.
func Min(stack []int) int {
	index := 0
	min := stack[0]
	for i := 0; i < len(stack); i++ {
		if stack[i] < min {
			min = stack[i]
			index = i
		}
	}
	return index
}

// Median calculates the median value of a slice of integers.
func Median(nums []int) int {
	slices.Sort(nums)
	if len(nums)%2 == 0 {
		return Round((float64(nums[len(nums)/2]) + float64(nums[(len(nums)/2)-1])) / 2.0)
	} else {
		return nums[len(nums)/2]
	}
}

// Round rounds a float64 to the nearest integer.
func Round(n float64) int {
	if int(n*10)%10 >= 5 {
		return int(math.Ceil(n))
	}
	return int(math.Floor(n))
}

// Pushswap implements the basic push swap algorithm for sorting stack A.
// It returns a string of instructions performed during the sorting process.
func Pushswap(stackA, stackB []int) (instruction string,validity bool) {
	instruction = ""
	for len(stackA) > 1 && !IsSorted(stackA) {
		index := Min(stackA)
		if index <= len(stackA)/2 {
			if index == 1 {
				SwapA(&stackA, &stackB)
				instruction += "sa\n"
			} else if index == 0 {
				PushB(&stackA, &stackB)
				instruction += "pb\n"
			}  else {
				RotateA(&stackA, &stackB)
				instruction += "ra\n"
			}
		} else {
			ReverseRotateA(&stackA, &stackB)
			instruction += "rra\n"
			index++
			if index == len(stackA)-1 {
				PushB(&stackA, &stackB)
				instruction += "pb\n"
			}
		}
	}
	for len(stackB) != 0 {
		PushA(&stackA, &stackB)
		instruction += "pa\n"
	}
	return instruction, IsSorted(stackA)
}

// Pushswapmax implements an optimized push swap algorithm for larger stacks.
// It uses the median value to partition the stack before sorting.
func Pushswapmax(stackA, stackB []int, m int) (instruction string,vaidity bool) {
	instruction = ""
	// Partition stackA around the median
	for len(stackA) > 1 {
		if stackA[0] <= m {
			PushB(&stackA, &stackB)
			instruction += "pb\n"
		} else {
			RotateA(&stackA, &stackB)
			instruction += "ra\n"
			if stackA[0] > m {
				ReverseRotateA(&stackA, &stackB)
				instruction += "rra\n"
				break
			}
		}
	}
	// Push elements back from stackB to stackA
	for len(stackB) > 0 {
		PushA(&stackA, &stackB)
		instruction += "pa\n"
	}
	return instruction, IsSorted(stackA)
}

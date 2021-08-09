package sort

import "testing"

type testpair struct {
	array       []int
	sortedArray []int
}

var testCases = []testpair{{[]int{3, 4, 1, 2, 5, 7, -1, 0}, []int{-1, 0, 1, 2, 3, 4, 5, 7}},
	{[]int{1, -5, 8, 21, 30, 15, -12, 10, 3}, []int{-12, -5, 1, 3, 8, 10, 15, 21, 30}},
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestQuickSort(t *testing.T) {
	for _, tests := range testCases {
		result := tests.array
		quickSort(result)
		if !Equal(result,tests.sortedArray) {
			t.Errorf("Have %d/nwant %d", result, tests.sortedArray)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tests := range testCases {
		result := tests.array
		bubbleSort(result)
		if !Equal(result,tests.sortedArray) {
			t.Errorf("Have %d/nwant %d", result, tests.sortedArray)
		}
	}
}

func TestSelectSort(t *testing.T){
	for _, tests := range testCases {
		result := tests.array
		selectSort(result)
		if !Equal(result,tests.sortedArray) {
			t.Errorf("Have %d/nwant %d", result, tests.sortedArray)
		}
	}
}

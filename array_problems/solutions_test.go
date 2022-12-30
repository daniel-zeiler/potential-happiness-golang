package array_problems

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	var result = [][]int{{1, 6}, {8, 10}, {15, 18}}
	var output = merge(intervals)
	if !reflect.DeepEqual(result, output) {
		t.Fatalf("not equal")
	}
	intervals = [][]int{{1, 4}, {0, 4}}
	result = [][]int{{0, 4}}
	output = merge(intervals)
	if !reflect.DeepEqual(result, output) {
		t.Fatalf("not equal")
	}
}

func TestArrayPairSum(t *testing.T) {
	var input = []int{1, 4, 3, 2}
	var result = 4
	if result != arrayPairSum(input) {
		t.Fatalf("not equal")
	}
	input = []int{6, 2, 6, 5, 1, 2}
	result = 9
	if result != arrayPairSum(input) {
		t.Fatalf("not equal")
	}
}

func TestSortArrayByParity(t *testing.T) {
	var input = []int{3, 1, 2, 4}
	var result = []int{2, 4, 3, 1}
	var output = SortArrayByParity(input)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("not equal")
	}
}

func TestReplaceElements(t *testing.T) {
	var input = []int{17, 18, 5, 4, 6, 1}
	var result = []int{18, 6, 6, 6, 1, -1}
	var output = replaceElements(input)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("not equal")
	}
}

func TestPermute(t *testing.T) {
	var input = []int{1, 2, 3}
	var result = [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	var output = permute(input)
	fmt.Printf("%v", output)
	if !reflect.DeepEqual(output, result) {
		t.Fatalf("not equal")
	}
}

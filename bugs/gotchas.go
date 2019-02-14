package main

import (
	"sort"
	"strconv"
)

// сюда вам надо писать функции, которых не хватает, чтобы проходили тесты в gotchas_test.go

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(data []int) string {
	var result string
	for _, v := range data {
		result += strconv.FormatInt(int64(v), 10)
	}
	return result
}

func MergeSlices(first []float32, second []int32) []int {
	var result []int
	for _, v := range first {
		result = append(result, int(v))
	}
	for _, v := range second {
		result = append(result, int(v))
	}
	return result
}

func GetMapValuesSortedByKey(m map[int]string) []string {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var result []string
	for _, k := range keys {
		result = append(result, m[k])
	}
	return result
}

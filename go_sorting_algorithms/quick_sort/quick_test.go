package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	tables := []struct {
		original []int
		sorted   []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{5, 8, 2, 9}, []int{2, 5, 8, 9}},
		{[]int{-9, -3, 9, -12}, []int{-12, -9, -3, 9}},
	}

	for _, table := range tables {
		array1 := table.original
		array2 := make([]int, len(array1))
		copy(array2, array1)
		shuffle(array1)
		QuickSort(array1, 0, len(array1)-1)
		sort.Ints(array2)

		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				t.Errorf("Sorted output for %v is incorrect; got: %v want: %v",
					table.original, array1, array2)
				break
			}
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrayRandom1 := make([]int, random.Intn(1000))
	for i := range arrayRandom1 {
		arrayRandom1[i] = random.Intn(10000)
	}
	arrayRandom2 := make([]int, len(arrayRandom1))
	copy(arrayRandom2, arrayRandom1)
	shuffle(arrayRandom1)
	QuickSort(arrayRandom1, 0, len(arrayRandom1)-1)
	sort.Ints(arrayRandom2)

	for i := 0; i < len(arrayRandom1); i++ {
		if arrayRandom1[i] != arrayRandom2[i] {
			t.Errorf("Quick Sort: %v \nSystem Sort: %v",
				arrayRandom1, arrayRandom2)
			break
		}
	}
}

func TestQuickSortOptimized(t *testing.T) {
	tables := []struct {
		original []int
		sorted   []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{5, 8, 2, 9}, []int{2, 5, 8, 9}},
		{[]int{-9, -3, 9, -12}, []int{-12, -9, -3, 9}},
	}

	for _, table := range tables {
		array1 := table.original
		array2 := make([]int, len(array1))
		copy(array2, array1)
		shuffle(array1)
		QuickSortOptimized(array1, 0, len(array1)-1)
		sort.Ints(array2)

		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				t.Errorf("Sorted output for %v is incorrect; got: %v want: %v",
					table.original, array1, array2)
				break
			}
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrayRandom1 := make([]int, random.Intn(1000))
	for i := range arrayRandom1 {
		arrayRandom1[i] = random.Intn(10000)
	}
	arrayRandom2 := make([]int, len(arrayRandom1))
	copy(arrayRandom2, arrayRandom1)
	shuffle(arrayRandom1)
	QuickSortOptimized(arrayRandom1, 0, len(arrayRandom1)-1)
	sort.Ints(arrayRandom2)

	for i := 0; i < len(arrayRandom1); i++ {
		if arrayRandom1[i] != arrayRandom2[i] {
			t.Errorf("Quick Sort Optimized: %v \nSystem Sort: %v",
				arrayRandom1, arrayRandom2)
			break
		}
	}
}

func TestQuick3WaySort(t *testing.T) {
	tables := []struct {
		original []int
		sorted   []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{5, 8, 2, 9}, []int{2, 5, 8, 9}},
		{[]int{-9, -3, 9, -12}, []int{-12, -9, -3, 9}},
	}

	for _, table := range tables {
		array1 := table.original
		array2 := make([]int, len(array1))
		copy(array2, array1)
		shuffle(array1)
		Quick3WaySort(array1, 0, len(array1)-1)
		sort.Ints(array2)

		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				t.Errorf("Sorted output for %v is incorrect; got: %v want: %v",
					table.original, array1, array2)
				break
			}
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrayRandom1 := make([]int, random.Intn(1000))
	for i := range arrayRandom1 {
		arrayRandom1[i] = random.Intn(10000)
	}
	arrayRandom2 := make([]int, len(arrayRandom1))
	copy(arrayRandom2, arrayRandom1)
	shuffle(arrayRandom1)
	Quick3WaySort(arrayRandom1, 0, len(arrayRandom1)-1)
	sort.Ints(arrayRandom2)

	for i := 0; i < len(arrayRandom1); i++ {
		if arrayRandom1[i] != arrayRandom2[i] {
			t.Errorf("Quick Sort Optimized: %v \nSystem Sort: %v",
				arrayRandom1, arrayRandom2)
			break
		}
	}
}

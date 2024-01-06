package sortingBenchmark

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
)

var arrLen int = 1000
var ar []int = make([]int, arrLen)

func init() {
	initArrByRandomValue(ar, 100)
	fmt.Printf("Массив из %d элементов.\n", arrLen)
}

func initArrByRandomValue(arr []int, maxValue int) {
	for key := range arr {
		arr[key] = rand.Intn(maxValue)
	}
}

func bubbleSort(arr []int) {
	for i := 0; i+1 < len(arr); i++ {
		for j := 0; j+1 < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func shakingSort(arr []int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		for i := left; i+1 < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

		for j := right; j > left; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		left++
		right--
	}
}

func combSort(arr []int) {
	const factor float64 = 1.224733
	//var step int = len(arr) - 1
	step := int(math.Floor(float64(len(arr)-1) / factor))
	for step >= 1 {
		for i := 0; i+step < len(arr); i++ {
			if arr[i] > arr[i+step] {
				arr[i], arr[i+step] = arr[i+step], arr[i]
			}
		}
		step = int(math.Floor(float64(step) / factor))
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > value; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = value
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		indexOfMinValue := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[indexOfMinValue] {
				indexOfMinValue = j
			}
		}
		if indexOfMinValue != i {
			arr[i], arr[indexOfMinValue] = arr[indexOfMinValue], arr[i]
		}
	}
}

func quickSort(arr []int) {
	quickSortImpl(arr, 0, len(arr)-1)
}

func quickSortImpl(arr []int, start int, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		quickSortImpl(arr, start, pivot-1)
		quickSortImpl(arr, pivot+1, end)
	}
}

func partition(arr []int, start int, end int) int {
	value := arr[end]
	pivot := start
	for i := start; i < end; i++ {
		if arr[i] <= value {
			arr[i], arr[pivot] = arr[pivot], arr[i]
			pivot++
		}
	}
	arr[pivot], arr[end] = arr[end], arr[pivot]
	return pivot
}

func quickSortByCenter(arr []int) {
	quickSortByCenterImpl(arr, 0, len(arr)-1)
}

func quickSortByCenterImpl(arr []int, start int, end int) {
	left := start
	right := end
	center := arr[(left+right)/2]
	for left <= right {
		for arr[right] > center {
			right--
		}
		for arr[left] < center {
			left++
		}
		if left <= right {
			arr[right], arr[left] = arr[left], arr[right]
			left++
			right--
		}
	}
	if right > start {
		quickSortByCenterImpl(arr, start, right)
	}
	if left < end {
		quickSortByCenterImpl(arr, left, end)
	}
}

// Сортировка пузырьком
func BenchmarkBubbleSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		bubbleSort(list)
	}
}

// Сортировка перемешиванием
func BenchmarkShakingSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		shakingSort(list)
	}
}

// Сортировка расческой
func BenchmarkCombSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		combSort(list)
	}
}

// Сортировка вставкой
func BenchmarkInsertionSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		insertionSort(list)
	}
}

// Сортировка выбором
func BenchmarkSelectionSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		selectionSort(list)
	}
}

// Быстрая сортировка. Опорный элемент последний.
func BenchmarkQuickSortByEndElement(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		quickSort(list)
	}
}

// Быстрая сортировка. Опорный элемент в центре.
func BenchmarkQuickSortByCenterElement(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		quickSortByCenter(list)
	}
}

// Сортировка в Go sort.Slice().
func BenchmarkGoSortSliceSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		sort.Slice(list, func(i, j int) bool {
			return list[i] < list[j]
		})
	}
}

// Сортировка в Go sort.SliceStable().
func BenchmarkGoSortSliceStableSort(b *testing.B) {
	var list []int = make([]int, arrLen)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, ar)
		b.StartTimer()
		sort.SliceStable(list, func(i, j int) bool {
			return list[i] < list[j]
		})
	}
}

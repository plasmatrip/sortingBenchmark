package sortingbenchmark

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

// var lenArr int
var arr []int
var arr1 []int
var arr2 []int
var arr3 []int
var arr4 []int
var arr5 []int
var arr6 []int
var arr7 []int
var arr8 []int
var arr9 []int

func init() {
	lenArr := 10
	arr := make([]int, lenArr)
	arr1 := make([]int, lenArr)
	arr2 := make([]int, lenArr)
	arr3 := make([]int, lenArr)
	arr4 := make([]int, lenArr)
	arr5 := make([]int, lenArr)
	arr6 := make([]int, lenArr)
	arr7 := make([]int, lenArr)
	arr8 := make([]int, lenArr)
	arr9 := make([]int, lenArr)
	initArrByRandomValue(arr, 100)
	copy(arr1, arr)
	copy(arr2, arr)
	copy(arr3, arr)
	copy(arr4, arr)
	copy(arr5, arr)
	copy(arr6, arr)
	copy(arr7, arr)
	copy(arr8, arr)
	copy(arr9, arr)

	fmt.Printf("Массив из %d элементов.\n", lenArr)
	fmt.Println(arr)
	fmt.Println(arr1)
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

func BenchmarkBuubleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort(arr1)
	}
}

// fmt.Print("Сортирвка пузырьком. ")
// startTime = time.Now()
// bubbleSort(arr1)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr1)

// fmt.Print("Сортирвка перемешиванием. ")
// startTime = time.Now()
// shakingSort(arr2)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr2)

// fmt.Print("Сортирвка расческой. ")
// startTime = time.Now()
// combSort(arr3)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr3)

// fmt.Print("Сортирвка вставкой. ")
// startTime = time.Now()
// insertionSort(arr4)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr4)

// fmt.Print("Сортирвка выбором. ")
// startTime = time.Now()
// selectionSort(arr5)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr5)

// fmt.Print("Быстрая сортировка. Опорный элемент последний. ")
// startTime = time.Now()
// quickSort(arr6)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr6)

// fmt.Print("Быстрая сортировка. Опорный элемент в центре. ")
// startTime = time.Now()
// quickSortByCenter(arr7)
// fmt.Println("Время сортировки: ", time.Since(startTime))
// //printArr(arr7)

// fmt.Print("Сортировка в Go sort.Slice(). ")
// startTime = time.Now()
// sort.Slice(arr8, func(i, j int) bool {
// 	return arr8[i] < arr8[j]
// })
// fmt.Println("Время сортировки: ", time.Since(startTime))

// fmt.Print("Сортировка в Go sort.SliceStable(). ")
// startTime = time.Now()
// sort.SliceStable(arr9, func(i, j int) bool {
// 	return arr9[i] < arr9[j]
// })
// fmt.Println("Время сортировки: ", time.Since(startTime))

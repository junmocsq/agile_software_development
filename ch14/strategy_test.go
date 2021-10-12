package ch14

import (
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	arr := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		arr[i] = rand.Intn(1000)
	}
	arr2 := make([]int, 20)
	t.Log("old:", arr)
	copy(arr2, arr)
	intArr1 := NewIntComparer(arr)
	Sort(NewBubbleSort(intArr1))
	t.Log("bubble    sort:", arr)

	intArr2 := NewIntComparer(arr2)
	Sort(NewInsertionSort(intArr2))
	t.Log("insertion sort:", arr2)

	sarr := make([]string, 20)
	sarr2 := make([]string, 20)
	for i := 0; i < 20; i++ {
		temp := []byte{}
		for j := 0; j <= rand.Intn(10); j++ {
			temp = append(temp, byte(rand.Intn(26))+'a')
		}
		sarr[i] = string(temp)
	}
	copy(sarr2, sarr)
	t.Log("old string:", sarr)

	stringArr1 := NewStringComparer(sarr)
	Sort(NewBubbleSort(stringArr1))
	t.Log("bubble    sort:", sarr)

	stringArr2 := NewStringComparer(sarr2)
	Sort(NewInsertionSort(stringArr2))
	t.Log("insertion sort:", sarr2)
}

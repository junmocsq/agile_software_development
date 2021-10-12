package ch14

type valueComparer interface {
	compare(i, k int) int
	length() int
	swap(i, k int)
	value(i int) interface{}
}

type intComparer struct {
	arr []int
}

func NewIntComparer(arr []int) *intComparer {
	return &intComparer{
		arr: arr,
	}
}
func (s *intComparer) compare(i, j int) int {
	return s.value(i).(int) - s.value(j).(int)
}
func (s *intComparer) length() int {
	return len(s.arr)
}
func (s *intComparer) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}
func (s *intComparer) value(i int) interface{} {
	return s.arr[i]
}

type stringComparer struct {
	arr []string
}

func NewStringComparer(arr []string) *stringComparer {
	return &stringComparer{
		arr: arr,
	}
}
func (s *stringComparer) compare(i, j int) int {
	if s.value(i).(string) > s.value(j).(string) {
		return 1
	} else {
		return -1
	}
}
func (s *stringComparer) length() int {
	return len(s.arr)
}
func (s *stringComparer) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}
func (s *stringComparer) value(i int) interface{} {
	return s.arr[i]
}

func Sort(s Sorter) {
	s.Sort()
}

type Sorter interface {
	Sort()
}
type bubbleSort struct {
	s valueComparer
}

func NewBubbleSort(s valueComparer) Sorter {
	return &bubbleSort{s: s}
}

func (s *bubbleSort) Sort() {
	for i := s.s.length() - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s.s.compare(j, j+1) > 0 {
				s.s.swap(j, j+1)
			}
		}
	}
}

type insertionSort struct {
	s valueComparer
}

func NewInsertionSort(s valueComparer) Sorter {
	return &insertionSort{s: s}
}
func (s *insertionSort) Sort() {
	for i := 1; i < s.s.length(); i++ {

		for j := i - 1; j >= 0; j-- {
			if s.s.compare(j, j+1) > 0 {
				s.s.swap(j, j+1)
			} else {
				break
			}
		}
	}
}

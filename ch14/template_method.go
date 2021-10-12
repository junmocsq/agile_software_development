package ch14

import "fmt"

type app interface {
	Init()
	Idle()
	Cleanup()
	SetDone()
	Done() bool
	Run(obj app)
}

type abstractApp struct {
	isDone bool
}

func (a *abstractApp) SetDone() {
	a.isDone = true
}

func (a *abstractApp) Done() bool {
	return a.isDone
}

func (a *abstractApp) Run(obj app) {
	obj.Init()
	for !obj.Done() {
		obj.Idle()
	}
	obj.Cleanup()
}

type Ftoc struct {
	abstractApp
	arr []int
}

func NewFtoc() app {
	return &Ftoc{}
}

func (f *Ftoc) Init() {
	f.arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
}

func (f *Ftoc) Idle() {
	if len(f.arr) == 0 {
		f.SetDone()
		return
	}
	fmt.Println("ftoc idle:", f.arr[0])
	f.arr = f.arr[1:]
}

func (f *Ftoc) Cleanup() {
	fmt.Println("cleanup")
}

type MMY struct {
	abstractApp
	arr []string
}

func NewMMY() app {
	return &MMY{}
}

func (f *MMY) Init() {
	f.arr = []string{"unix", "macos", "solaris", "windows", "dos", "os360"}
}

func (f *MMY) Idle() {
	if len(f.arr) == 0 {
		f.SetDone()
		return
	}
	fmt.Println("MMY idle:", f.arr[0])
	f.arr = f.arr[1:]
}

func (f *MMY) Cleanup() {
	fmt.Println("MMY cleanup")
}

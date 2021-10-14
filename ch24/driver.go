package ch24

import "fmt"

type TimeSource interface {
	SetObserver(observer ClockObserver)
}
type MockTimeSource struct {
	itsObservers []ClockObserver
}

func NewMockTimeSource() *MockTimeSource {
	return &MockTimeSource{}
}
func (mt *MockTimeSource) SetTime(hours, minutes, seconds int) {
	for _, observer := range mt.itsObservers {
		observer.Update(hours, minutes, seconds)
	}
}
func (mt *MockTimeSource) RegisterObserver(d ClockObserver) {
	mt.itsObservers = append(mt.itsObservers, d)
}

type TimeSink interface {
	SetTime(hours, minutes, seconds int)
}

type MockTimeSink struct {
	itsHours, itsMinutes, itsSeconds int
}

func NewMockTimeSink() *MockTimeSink {
	return &MockTimeSink{}
}

func (mts *MockTimeSink) Update(hours, minutes, seconds int) {
	mts.itsHours = hours
	mts.itsMinutes = minutes
	mts.itsSeconds = seconds
	mts.PrintTime()
}

func (mts *MockTimeSink) PrintTime() {
	fmt.Printf("%02d:%02d:%02d\n", mts.itsHours, mts.itsMinutes, mts.itsSeconds)
}

type ClockObserver interface {
	Update(hours, minutes, seconds int)
}

//type ClockDriver struct {
//	itsSinks []TimeSink
//}
//
//func NewClockDriver(source TimeSource) *ClockDriver {
//	c:= &ClockDriver{}
//	source.SetObserver(c)
//	return c
//}
//
//func (c *ClockDriver) Update(hours, minutes, seconds int) {
//	for _,sink := range c.itsSinks{
//		sink.SetTime(hours, minutes, seconds)
//	}
//}
//
//func (c *ClockDriver)addTimeSink(sink TimeSink)  {
//	c.itsSinks = append(c.itsSinks,sink)
//}
